(ns aupro.user
  (:require [centipair.ui :as ui]
            [centipair.components.input :as input]
            [reagent.core :as r]
            [centipair.ajax :as ajax]))

(def email (r/atom {:id "email" :type "email" :class "cfi" :placeholder "Enter Email" :label "Email"}))
(def password (r/atom {:id "password" :type "password" :class "cfi" :placeholder "Enter Password" :label "Password"}))
(def phone (r/atom {:id "phone" :type "text" :class "cfi" :placeholder "Enter Phone"}))
(def full-name (r/atom {:id "full-name" :type "text" :class "cfi" :placeholder "Enter Full Name"}))

(def no-account-link (r/atom {:text "No Account? " :label "Sign up" :href "#/register"}))
(def forgot-password-link (r/atom {:text "Forgot password? " :label "Reset password" :href "#/reset-password"}))
(def already-registered-link (r/atom {:text "Already registered? " :label "Login" :href "#/login"}))

(def login
  [])
(def login-button (r/atom {:label "Login" :on-click login}))


(defn register
  []
  (ajax/recap-form-post "/api/uauth/register" [email password full-name phone] 
                  (fn [response])))

(def register-button (r/atom {:label "Register" :on-click register}))

(defn text [field]
  [:div
   [:label {:for (:id @field), :class "sr-only"} (:label @field)]
   [:div {:class "relative"}
    (input/text field)
    [:span {:class "cfps"} (:icon @field)]]
   [:p {:id (str (:id @field) "-message") :class "link-error"} (:message @field)]])

(defn footer-link
  [link]
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
      [:button {:type "input", :class "btn btn-primary w-full" :on-click (:on-click @button)} (:label @button)]
      (map footer-link footer-links)]]])


(defn login-page 
  []
  (generate-form "Login"
                 "Sign in to your account"
                 [email password]
                 login-button
                 [no-account-link forgot-password-link]))

(defn render-login
  []
  (ui/render-ui login-page "app"))

(defn register-page
  []
  (generate-form "Sign up"
                 "Register a new account"
                 [full-name email password phone]
                 register-button
                 [already-registered-link]))

(defn render-register
  []
  (ui/render-ui register-page "app"))

(defn reset-password-page
  []
  (generate-form "Reset Password"
                 "Enter email to reset password"
                 [email]
                 login-button
                 [no-account-link]))


(defn render-reset-password
  []
  (ui/render-ui reset-password-page "app"))