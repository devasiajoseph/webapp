(ns aupro.form
  (:require [centipair.components.input :as input]))




(defn text [field]
  ^{:key (:id @field)}
  [:div
   [:label {:for (:id @field), :class "sr-only"} (:label @field)]
   [:div {:class "relative"}
    (input/text field)
    [:span {:class "cfps"} (:icon @field)]]
   [:p {:id (str (:id @field) "-message") :class "link-error"} (:message @field)]])

(defn footer-link
  [link]
  ^{:key (:id @link)}
  [:p {:class "text-center text-sm text-gray-500"} (:text @link)
   [:a {:class "underline", :href (:href @link)} (:label @link)]])


(defn generate-form
  [title header inputs button footer-links]
  [:div {:class "cfc"}
   [:div {:class "cf card bg-base-100"}
    [:h1 {:class "cfh"} title]
    [:form {:class "cff"}
     [:p {:class "text-center text-lg font-medium"} header]
     (doall (map text inputs))
     [:a {:type "input", :class "btn btn-primary w-full" :on-click (:on-click @button)} (:label @button)]
     (doall (map footer-link footer-links))]]])

