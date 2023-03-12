(ns aupro.dashboard
  (:require [centipair.ui :as ui]))



(defn dashboard-view
  []
  [:div {:id "dashboard-container"}
   [:div {:class "flex flex-col md:flex-row"}
    [:div {:class ""} [:a {:href "#/profile/list"}]]]])



(defn render-dashboard
  []
  (ui/render-ui dashboard-view "app"))

