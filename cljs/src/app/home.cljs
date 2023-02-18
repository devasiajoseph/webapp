(ns app.home
   (:require [reagent.core :as r]
             [centipair.components.input :as input]
              [centipair.components.notifier :as notifier]
             [centipair.ui :as ui]
             [centipair.ajax :as ajax]))

(defonce email (r/atom {:type "text" :id "email" :class "form-control"}))

(def search-box (r/atom {:type "text" :value "" :id "search-input" :placeholder "Enter bitcoin address"
                         :show-results false}))

(def bitquery (r/atom {:balance 0}))


(defn search-bitcoin
  []
  (ajax/get-json "/api/bitcoin/balance" {:addr (:value @search-box)}
                 (fn [response]
                   
                   (do
                     (notifier/notify 200)
                     (reset! bitquery response)
                     (swap! search-box assoc :show-results true)))))

(defn home-page []
  [:div {:class "mx-auto content-center max-w-2xl"}
   [:div {:id "home-search-box" }
    [:h3 {:class "flex justify-center font-bold text-2xl"} "Search Bitcoin"]
    [:label {:for "default-search", :class "mb-2 text-sm font-medium text-gray-900 sr-only"} "Search"]
    [:div {:class "relative"}
     [:div {:class "absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none"}
      [:svg {:aria-hidden "true", :class "w-5 h-5 text-gray-500 ", :fill "none", :stroke "currentColor", :viewBox "0 0 24 24", :xmlns "http://www.w3.org/2000/svg"}
       [:path {:stroke-linecap "round", :stroke-linejoin "round", :stroke-width "2", :d "M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"}]]]
     [:input {:type (:type @search-box), :id (:id @search-box), :placeholder (:placeholder @search-box) :value (:value @search-box)
              :on-change #(input/update-value search-box (-> % .-target .-value))
              :class "search-box"}]
     [:button {:type "submit", :class "search-btn" :on-click search-bitcoin} "Search"]]
    [:div {:class "mt-6"}
     [:div {:class (str "card" (if (:show-results @search-box) "" " hidden"))}
      [:div {:class "px-6 py-6"}
       [:div {:class "text-sm text-gray-400 mb-2"} "Bitcoin Balance"]

       [:p [:span {:class "text-gray-700 text-base text-5xl" } (:balance @bitquery)]
        [:span {:class "text-sm text-gray-500" } " Satoshis"]]
       [:div {:class "mb-5 border-b"}]
       [:table {:class "w-full"}
        [:tbody 
         [:tr [:td {:class "font-bold"}"Total Received:"] [:td (:total_received @bitquery)]]
         [:tr [:td "Total Sent:"] [:td (:total_sent @bitquery)]]
         [:tr [:td "Unconfirmed Balance:"] [:td (:unconfirmed_balance @bitquery)]]
         [:tr [:td "Final Balance:"] [:td (:final_balance @bitquery)]]]]]]]]])

;;
(defn render-home []
  (ui/render-ui home-page "app"))

(defn about-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to about"]
   [:p "This is about page"]])


(defn render-about []
  (ui/render-ui about-page "app"))


