(ns aupro.dashboard
  (:require [centipair.ui :as ui]))



(defn dashboard-view
  []
  [:div {:id "dashboard-container"}
   [:div {:class "flex flex-col md:flex-row"}
    [:div {:class "shadow-xl card p-5 bg-white"}
     [:img {:src "/static/icons/user.png" :class "mx-auto h-10 w-10"}]
     [:div {:class "card-body"}
      [:a {:href "#/profile/list/1" :class ""} "Profiles"]]]]])



(defn render-dashboard
  []
  (ui/render-ui dashboard-view "app"))

