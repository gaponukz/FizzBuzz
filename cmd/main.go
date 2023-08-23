package main

import (
	"fizzbuzz/src/application/rules"
	"fizzbuzz/src/application/strategies"
	"fizzbuzz/src/infrastructure/controller"
	"fizzbuzz/src/infrastructure/presentation"
	"net/http"
)

func main() {
	fizzCondition := strategies.NewDivStrategy(3)
	buzzCondition := strategies.NewDivStrategy(5)
	service := rules.NewRulesCollection(
		rules.NewRule("FizzBuzz", strategies.NewAndStrategy(fizzCondition, buzzCondition)),
		rules.NewRule("Fizz", fizzCondition),
		rules.NewRule("Buzz", buzzCondition),
	)

	web := controller.NewController(service, presentation.NewRawPresenter())
	httpRoute := http.NewServeMux()

	httpRoute.HandleFunc("/run", web.Convert)
	server := &http.Server{
		Addr:    ":8080",
		Handler: httpRoute,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
