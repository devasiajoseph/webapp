(ns app.uauth
  (:require [reagent.core :as r]
            [centipair.ui :as ui]
            [centipair.components.input :as input]
            [centipair.components.form :as form]))

(def userid (r/atom {:id "userid" :type "text" :class "form-control"}))
(def password (r/atom {:id "password" :type "password" :class "form-control"}))


(defn login-form 
  []
  [:div {:class "form-data"}
   [:div {:class "forms-inputs mb4"} [:span "Phone"]
    (input/text userid)
    [:label {:id "userid-error"}]]
   [:div {:class "forms-inputs mb4"} [:span "Password"]
    (input/text password)
    [:label {:id "password-error"}]]
   [:div {:class "mb-3"} [:a {:class "btn btn-primary w-100"} "Login"]]
   ])

(defn login-page
  []
  (form/card-form 6 login-form))


(defn render-login-form
  []
  (ui/render-ui login-page "app"))