(ns app.home
   (:require [reagent.core :as r]
             [centipair.components.input :as input]
             [centipair.ui :as ui]))

(defonce email (r/atom {:type "text" :id "email" :class "form-control"}))

(defn home-page []
  [:div {:id "home-search-box"}
   [:div {:class "flex justify-start font-bold text-xl"} "Enter bitcoin address to search"]
   [:label {:for "default-search", :class "mb-2 text-sm font-medium text-gray-900 sr-only"} "Search"]
   [:div {:class "relative"}
    [:div {:class "absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none"}
     [:svg {:aria-hidden "true", :class "w-5 h-5 text-gray-500 ", :fill "none", :stroke "currentColor", :viewBox "0 0 24 24", :xmlns "http://www.w3.org/2000/svg"}
      [:path {:stroke-linecap "round", :stroke-linejoin "round", :stroke-width "2", :d "M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"}]]]
    [:input {:type "search", :id "default-search", :class "block w-full p-4 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 ", :placeholder "Enter bitcoin address"}]
    [:button {:type "submit", :class "text-white absolute right-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 "} "Search"]]])


(defn render-home []
  (ui/render-ui home-page "app"))

(defn about-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to about"]
   [:p "This is about page"]])


(defn render-about []
  (ui/render-ui about-page "app"))


