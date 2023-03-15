(ns centipair.location
  (:require [reagent.core :as reagent]
            [centipair.components.input :as cin]
            [centipair.db :as db]))



(def country (reagent/atom {:id "country_id" :type "select" :value "IN" :label "Country"}))

