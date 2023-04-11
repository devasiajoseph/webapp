(ns aupro.dashboard
  (:require [centipair.ui :as ui]
            [reagent.core :as r]
            [centipair.spa :as spa]))




(def profiles (r/atom {:icon "address-card" :label "Profiles" :link "#/profile/list/1"}))
(def managers (r/atom {:icon "user" :label "Managers" :link "#/manager/list/1"}))

(defn icon
  [field]
  [:div {:class "flex flex-col items-center cursor-pointer" :on-click #(spa/redirect (:link @field))}
   [:i {:class (str "fa-solid fa-3x fa-" (:icon @field))}]
   [:span {:class "mt-2"} (:label @field)]])

(defn dashboard-view
  []
  [:div {:id "dashboard-container"}
   [:div {:class "grid grid-cols-4 gap-4"}
    (icon profiles)
    (icon managers)]])

(defn render-dashboard
  []
  (ui/render-ui dashboard-view "app"))

