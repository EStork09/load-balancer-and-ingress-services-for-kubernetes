/*
* [2013] - [2019] Avi Networks Incorporated
* All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package integrationtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/onsi/gomega"
	"gitlab.eng.vmware.com/orion/akc/pkg/cache"
	"gitlab.eng.vmware.com/orion/akc/pkg/k8s"
)

func TestCacheGETOKStatus(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if strings.Contains(r.URL.EscapedPath(), "virtualservice") {
			data, _ := ioutil.ReadFile("avimockobjects/shared_vs_mock.json")

			fmt.Fprintln(w, string(data))
		} else if strings.Contains(r.URL.EscapedPath(), "poolgroup") {
			data, _ := ioutil.ReadFile("avimockobjects/poolgroups_mock.json")

			fmt.Fprintln(w, string(data))
		} else if strings.Contains(r.URL.EscapedPath(), "pool") {
			data, _ := ioutil.ReadFile("avimockobjects/pool_mock.json")

			fmt.Fprintln(w, string(data))
		} else if strings.Contains(r.URL.EscapedPath(), "vsdatascript") {
			data, _ := ioutil.ReadFile("avimockobjects/datascript_http_mock.json")
			fmt.Fprintln(w, string(data))
		} else {
			// This is used for /login --> first request to controller
			fmt.Fprintln(w, string(`{"dummy" :"data"}`))
		}

	}))
	defer ts.Close()
	url := strings.Split(ts.URL, "https://")[1]
	os.Setenv("CTRL_USERNAME", "admin")
	os.Setenv("CTRL_PASSWORD", "admin")
	os.Setenv("CTRL_IPADDRESS", url)
	k8s.PopulateCache()
	// Verify the cache.
	cacheobj := cache.SharedAviObjCache()
	vsKey := cache.NamespaceName{Namespace: "admin", Name: "Shard-VS-5"}
	vs_cache, found := cacheobj.VsCache.AviCacheGet(vsKey)

	if !found {
		t.Fatalf("Cache not found for VS: %v", vsKey)
	} else {
		vs_cache_obj, ok := vs_cache.(*cache.AviVsCache)
		if !ok {
			t.Fatalf("Invalid VS object. Cannot cast.")
		}
		g.Expect(len(vs_cache_obj.PoolKeyCollection)).To(gomega.Equal(3))
		g.Expect(len(vs_cache_obj.PGKeyCollection)).To(gomega.Equal(1))
		g.Expect(len(vs_cache_obj.DSKeyCollection)).To(gomega.Equal(1))
	}
}

func TestCacheGETControllerUnavailable(t *testing.T) {
	ctrlUnavail := true
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ctrlUnavail {
			w.WriteHeader(http.StatusServiceUnavailable)
			ctrlUnavail = false
		} else {
			fmt.Fprintln(w, string(`{"dummy" :"data"}`))
		}

	}))
	defer ts.Close()
	url := strings.Split(ts.URL, "https://")[1]
	os.Setenv("CTRL_USERNAME", "admin")
	os.Setenv("CTRL_PASSWORD", "admin")
	os.Setenv("CTRL_IPADDRESS", url)
	k8s.PopulateCache()
	// Verify the cache.
	cacheobj := cache.SharedAviObjCache()
	vsKey := cache.NamespaceName{Namespace: "admin", Name: "Shard-VS-5"}
	_, found := cacheobj.VsCache.AviCacheGet(vsKey)
	if !found {
		// The older cache member should be available.
		t.Fatalf("Cache not found for VS: %v", vsKey)
	}
	vsKey = cache.NamespaceName{Namespace: "admin", Name: "Shard-VS-4"}
	_, found = cacheobj.VsCache.AviCacheGet(vsKey)
	if found {
		// The older cache member should be available.
		t.Fatalf("Cache found for VS: %v", vsKey)
	}
}

func TestCacheGETDependentObjectUnavailable(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	// Verify the state of the cache

	cacheobj := cache.SharedAviObjCache()
	vsKey := cache.NamespaceName{Namespace: "admin", Name: "Shard-VS-5"}
	vs_cache, found := cacheobj.VsCache.AviCacheGet(vsKey)
	if !found {
		// The older cache member should be available.
		t.Fatalf("Cache not found for VS: %v", vsKey)
	} else {
		vs_cache_obj, ok := vs_cache.(*cache.AviVsCache)
		if !ok {
			t.Fatalf("Invalid VS object. Cannot cast.")
		}
		g.Expect(len(vs_cache_obj.PoolKeyCollection)).To(gomega.Equal(3))
		g.Expect(len(vs_cache_obj.PGKeyCollection)).To(gomega.Equal(1))
		g.Expect(len(vs_cache_obj.DSKeyCollection)).To(gomega.Equal(1))
	}
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if strings.Contains(r.URL.EscapedPath(), "virtualservice") {
			data, _ := ioutil.ReadFile("avimockobjects/shared_vs_mock.json")
			fmt.Fprintln(w, string(data))
		} else if strings.Contains(r.URL.EscapedPath(), "poolgroup") {
			w.WriteHeader(http.StatusInternalServerError)
		} else if strings.Contains(r.URL.EscapedPath(), "pool") {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else if strings.Contains(r.URL.EscapedPath(), "vsdatascript") {
			data, _ := ioutil.ReadFile("avimockobjects/datascript_http_mock.json")
			fmt.Fprintln(w, string(data))
		} else {
			// This is used for /login --> first request to controller
			fmt.Fprintln(w, string(`{"dummy" :"data"}`))
		}
	}))
	defer ts.Close()
	url := strings.Split(ts.URL, "https://")[1]
	os.Setenv("CTRL_USERNAME", "admin")
	os.Setenv("CTRL_PASSWORD", "admin")
	os.Setenv("CTRL_IPADDRESS", url)
	k8s.PopulateCache()
	// Verify the cache.
	vs_cache, found = cacheobj.VsCache.AviCacheGet(vsKey)
	if !found {
		// The older cache member should be available.
		t.Fatalf("Cache not found for VS: %v", vsKey)
	} else {
		vs_cache_obj, ok := vs_cache.(*cache.AviVsCache)
		if !ok {
			t.Fatalf("Invalid VS object. Cannot cast.")
		}
		g.Expect(len(vs_cache_obj.PoolKeyCollection)).To(gomega.Equal(3))
		// The PG had a problem in GET operation, but we will retain the cache.
		g.Expect(len(vs_cache_obj.PGKeyCollection)).To(gomega.Equal(1))
		g.Expect(len(vs_cache_obj.DSKeyCollection)).To(gomega.Equal(1))
	}
	vsKey = cache.NamespaceName{Namespace: "admin", Name: "Shard-VS-4"}
	_, found = cacheobj.VsCache.AviCacheGet(vsKey)
	if found {
		t.Fatalf("Cache found for VS: %v", vsKey)
	}
}
