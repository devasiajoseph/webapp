(ns centipair.location
  (:require [reagent.core :as reagent]
            [centipair.location :as location]))



(def country (reagent/atom {:id "country_id" :type "select" :value "IN" :label "Country"}))

(defn get-country-list
  []
  )