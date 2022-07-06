(ns fintech.asset
  (:require [reagent.core :as r]
            [centipair.ui :as ui]))




(def asset-data (r/atom {:name "Bitcoin" :symbol "btc" :price "$20,184.81"}))

(defn price
  [])


(defn container
  []
  [:div 
   [:div {:class "row mb-5"}
    [:div {:class "col-md-2"}
     [:h2 (:name @asset-data)]
     [:span {:class "badge bg-light text-primary"} "BTC"]
     [:span {:class "badge bg-light text-secondary"} "cryptocurrency"]]
    [:div {:class "col"} [:h3 (:price @asset-data)]]]
   [:div {:class "row"}
    [:div {:class "col-md-2"}
     [:h2 (:name @asset-data)]
     [:span {:class "badge bg-light text-primary"} "BTC"]
     [:span {:class "badge bg-light text-secondary"} "cryptocurrency"]]
    [:div {:class "col"} [:h3 (:price @asset-data)]]]])

(defn render-asset-page 
  []
  (ui/render-ui container "app"))