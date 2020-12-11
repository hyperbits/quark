package quark

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) SetupRouter() *mux.Router {
	a.Router = mux.NewRouter().StrictSlash(true)
	return a.Router
}
func (a *App) SPA(path, dir string) {
	fileServer := http.FileServer(spaFileSystem{http.Dir(dir)})
	a.Router.PathPrefix(path).Handler(fileServer)
}
func (a *App) Static(path, dir string) {
	a.Router.PathPrefix(path).Handler(http.StripPrefix(dir, http.FileServer(http.Dir(dir))))
}
func (a *App) Get(route string, handler http.Handler) {
	a.Route(route, "GET", handler)
}
func (a *App) Post(route string, handler http.Handler) {
	a.Route(route, "POST", handler)
}
func (a *App) Put(route string, handler http.Handler) {
	a.Route(route, "PUT", handler)
}
func (a *App) Delete(route string, handler http.Handler) {
	a.Route(route, "DELETE", handler)
}
func (a *App) Route(route, method string, handler http.Handler) {
	a.Router.Handle(route, handler).Methods(method)
}

type spaFileSystem struct {
	fs http.FileSystem
}

func (nfs spaFileSystem) Open(path string) (http.File, error) {

	f, err := nfs.fs.Open(path)
	if err != nil {
		if f, err = nfs.fs.Open("index.html"); err != nil {
			return nil, err
		}
	}

	return f, nil
}
