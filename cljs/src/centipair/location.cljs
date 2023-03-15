(ns centipair.location
  (:require [reagent.core :as reagent]
            [centipair.ajax :as ajax]))



(def country (reagent/atom {:id "country_id" :type "select" :value "IN" :label "Country"}))

(defn get-country-list
  []
  )