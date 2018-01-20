package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
	"flag"
	"context"
	"os"
	"os/signal"
	"github.com/guesswho311/mux"
	"github.com/guesswho311/godip/variants"

	gae "github.com/guesswho311/godip/gae/common"
)

func corsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func preflight(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
}

func resolve(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	variantName := mux.Vars(r)["variant"]
	variant, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	p := &gae.Phase{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	state, err := p.State(variant)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err = state.Next(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Load the new godip phase from the state
	nextPhase := gae.NewPhase(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(nextPhase); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	variantName := mux.Vars(r)["variant"]
	variant, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	state, err := variant.Start()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	phase := gae.NewPhase(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(phase); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func listVariants(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(variants.Variants); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

/*func init() {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(preflight)
	variants := r.Path("/{variant}").Subrouter()
	variants.Methods("POST").HandlerFunc(resolve)
	variants.Methods("GET").HandlerFunc(start)
	r.Path("/").HandlerFunc(listVariants)
	http.Handle("/", r)
}*/

func main() {
	    var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(preflight)
	variants := r.Path("/{variant}").Subrouter()
	variants.Methods("POST").HandlerFunc(resolve)
	variants.Methods("GET").HandlerFunc(start)
	r.Path("/").HandlerFunc(listVariants)
	http.Handle("/", r)
	log.Println("test")
	log.Println(r)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println("shutting down")
    os.Exit(0)
	
}
