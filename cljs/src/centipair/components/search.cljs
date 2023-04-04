(ns centipair.components.search
  (:require [centipair.components.input :as input]))



(defn search-box 
  [field search-fn]
   [:div {:id "home-search-box"}
    [:label {:for "default-search", :class "mb-2 text-sm font-medium text-gray-900 sr-only"} "Search"]
    [:div {:class "relative"}
     [:div {:class "absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none"}
      [:svg {:aria-hidden "true", :class "w-5 h-5 text-gray-500 ", :fill "none", :stroke "currentColor", :viewBox "0 0 24 24", :xmlns "http://www.w3.org/2000/svg"}
       [:path {:stroke-linecap "round", :stroke-linejoin "round", :stroke-width "2", :d "M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"}]]]
     [:input {:type "text", :id (:id @field), :placeholder (:placeholder @field) :value (:value @field)
              :on-change #(input/update-value field (-> % .-target .-value))
              :class "search-box"}]
     [:button {:type "submit", :class "search-btn" :on-click search-fn} "Search"]]])