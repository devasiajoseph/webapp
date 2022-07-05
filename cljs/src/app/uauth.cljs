(ns app.uauth
  (:require [reagent.core :as r]
            [centipair.ui :as ui]
            [centipair.components.input :as input]
            [centipair.components.form :as form]
            [centipair.validators :as v]
            [centipair.ajax :as ajax]
            [app.api :as api]))

(def userid (r/atom {:id "userid" :type "text" :class "form-control form-control-lg" :validator v/required}))
(def password (r/atom {:id "password" :type "password" :class "form-control form-control-lg" :validator v/required}))

(defn login 
  []
  (ajax/form-post (api/url :user-login) [userid password] (fn [response])))

(def login-button (r/atom {:id "login-button" :label "Login" :on-click login}))

(defn login-form 
  []
  [:form
   [:div {:class " mb-4"} [:label "Phone"]
    (input/text userid)
    [:div {:id "userid-error" :class "invalid-field"} (:message @userid)]]
   [:div {:class "mb-4"} [:label "Password"]
    (input/text password)
    [:div {:id "userid-error" :class "invalid-field"} (:message @password)]]
   (input/button login-button [userid password])])

(defn login-page
  []
  (form/card-form 6 login-form))


(defn render-login-form
  []
  (ui/render-ui login-page "app"))