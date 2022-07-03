(ns app.home
   (:require [reagent.dom :as r]))


(defn home-page []
  [:div
   [:h2 "Welcome to APP"]
   [:p "This page is rendered from clojurescript reagent"]])


(defn render-home []
  (r/render [home-page] (.getElementById js/document "app")))

