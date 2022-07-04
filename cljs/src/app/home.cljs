(ns app.home
   (:require [reagent.core :as r]
             [centipair.components.input :as input]
             [centipair.ui :as ui]))

(defonce email (r/atom {:type "text" :id "email" :class "form-control"}))

(defn home-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to APP"]
   [:p "This page is rendered from clojurescript reagent"]
   [:p (input/text email)]])


(defn render-home []
  (ui/render-ui home-page "app"))

(defn about-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to about"]
   [:p "This is about page"]])


(defn render-about []
  (ui/render-ui about-page "app"))


