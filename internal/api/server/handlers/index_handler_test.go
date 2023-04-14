package handlers

//func TestIndexHandler(t *testing.T) {
//
//	r := chi.NewRouter()
//	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//		w.WriteHeader(http.StatusOK)
//	})
//
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	res, _ := TestRequest(t, ts, "GET", "/", nil)
//	defer res.Body.Close()
//
//	assert.Equal(t, res.StatusCode, http.StatusOK)
//}

/*
func TestHandler_IndexHandler(t *testing.T) {
	testTable := []struct {
		name        string
		server      *httptest.Server
		expectedErr error
	}{
		{
			name: "test index handler",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(http.StatusOK)
			})),
			expectedErr: nil,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, err := MakeHTTPCall(tc.server.URL)
			if err != tc.expectedErr {
				t.Error(err)
			}
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, http.StatusOK)
			assert.Equal(t, resp.Header.Get("Content-Type"), "text/html")
		})
	}
}
*/

/*func TestHandler_IndexHandler(t *testing.T) {

	responseRecorder := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("secretKey")

	configServer := configserver.ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		DebugLevel:      logrus.DebugLevel,
		StoreFile:       "/tmp/devops-metrics-db.json",
		Key:             "",
		DSN:             "",
		RootPath:        "file://./migrations",
		//TemplatePath:    "file://./web/templates/index.html",
	}

	srv := NewHandler(memStorage, nil, &configServer, nil, hashServer)

	router := chi.NewRouter()
	router.Get("/", srv.IndexHandler)

	var statusExpect = http.StatusOK
	var contentTypeExpect = "text/html"

	router.ServeHTTP(responseRecorder, httptest.NewRequest("GET", "/", nil))
	statusGet := responseRecorder.Code
	contentTypeGet := responseRecorder.Header().Get("Content-Type")

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, contentTypeExpect, contentTypeGet, fmt.Sprintf("Incorrect Content-Type. Expect %s, got %s", contentTypeExpect, contentTypeGet))

}
*/
