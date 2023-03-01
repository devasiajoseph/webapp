(ns aupro.dashboard
  (:require [centipair.ui :as ui]))



(defn dashboard-view
  []
  [:div {:id "dashboard-container"}
   [:div {:class "flex flex-col md:flex-row"}
    [:div {:class "card bg-base-100 shadow-xl flex-1 mx-2 my-5"}
     [:div {:class "card-body justify-center"}
      [:h2 {:class "card-title justify-center"} "My profile"]
      [:div {:class "card-actions justify-center"} "If a dog chews shoes whose shoes does he choose?"]
      [:div {:class "card-actions justify-center"}
       [:button {:class "btn btn-primary"} "Buy Now"]]]]
    [:div {:class "card bg-base-100 shadow-xl flex-1 mx-2 my-5"}
     [:div {:class "card-body justify-center"}
      [:h2 {:class "card-title justify-center"} "My profile"]
      [:p "If a dog chews shoes whose shoes does he choose?"]
      [:div {:class "card-actions justify-center"}
       [:button {:class "btn btn-primary"} "Buy Now"]]]]
    [:div {:class "card bg-base-100 shadow-xl flex-1 mx-2 my-5"}
     [:div {:class "card-body justify-center"}
      [:h2 {:class "card-title justify-center"} "My profile"]
      [:p "If a dog chews shoes whose shoes does he choose?"]
      [:div {:class "card-actions justify-center"}
       [:button {:class "btn btn-primary"} "Buy Now"]]]]
    [:div {:class "card bg-base-100 shadow-xl flex-1 mx-2 my-5"}
     [:div {:class "card-body"}
      [:h2 {:class "card-title"} "My profile"]
      [:p "If a dog chews shoes whose shoes does he choose?"]
      [:div {:class "card-actions justify-end"}
       [:button {:class "btn btn-primary"} "Buy Now"]]]]]])



(defn render-dashboard
  []
  (ui/render-ui dashboard-view "app"))

