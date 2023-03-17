(ns centipair.location
  (:require [reagent.core :as reagent]
            [centipair.ajax :as ajax]))



(def country (reagent/atom {:id "country_id" :type "select" :value "228" :label "Country"
                            :label-key "country_name" :value-key "country_id" :remote "/api/location/countries"
                            :remote-cache true}))