(ns aupro.user
  (:require [centipair.ui :as ui]))


(defn login-page 
  []
  [:div {:class "cfc"}
   [:div {:class "cf"}
    [:h1 {:class "cfh"} "Login"]
    [:form { :class "cff"}
     [:p {:class "text-center text-lg font-medium"} "Sign in to your account"]
     [:div
      [:label {:for "email", :class "sr-only"} "Email"]
      [:div {:class "relative"}
       [:input {:type "email", :class "cfi", :placeholder "Enter email"}]
       [:span {:class "cfps"} "✉"]]]
     [:div
      [:label {:for "password", :class "sr-only"} "Password"]
      [:div {:class "relative"}
       [:input {:type "password", :class "cfi", :placeholder "Enter password"}]
       [:span {:class "cfps" } "🔑"]]]
     [:button {:type "input", :class "cfb"} "Sign in"]
     [:p {:class "text-center text-sm text-gray-500"} "No account? "
      [:a {:class "underline", :href "#/register"} "Sign up"]]
     [:p {:class "text-center text-sm text-gray-500"} "Forgot password? "
      [:a {:class "underline", :href "#/reset-password"} "Reset password"]]
     ]]])


(defn render-login
  []
  (ui/render-ui login-page "app"))



(defn register-page
  []
  [:div {:class "cfc"}
   [:div {:class "cf"}
    [:h1 {:class "cfh"} "Sign up"]
    [:form {:class "cff"}
     [:p {:class "text-center text-lg font-medium"} "Register a new account"]
     [:div
      [:label {:for "email", :class "sr-only"} "Email"]
      [:div {:class "relative"}
       [:input {:type "email", :class "cfi", :placeholder "Enter email"}]
       [:span {:class "cfps"} "✉"]]]
     [:div
      [:label {:for "password", :class "sr-only"} "Password"]
      [:div {:class "relative"}
       [:input {:type "password", :class "cfi", :placeholder "Enter password"}]
       [:span {:class "cfps"}  "🔑"]]]
     [:button {:type "input", :class "cfb"} "Register"]
     [:p {:class "cfm"} "Already registered? "
      [:a {:class "underline", :href "#/login"} "Login"]]]]])


(defn render-register
  []
  (ui/render-ui register-page "app"))



(defn reset-password-page
  []
  [:div {:class "cfc"}
   [:div {:class "cf"}
    [:h1 {:class "cfh"} "Reset password"]
    [:form {:class "cff"}
     [:p {:class "text-center text-lg font-medium"} "Enter your email to reset password"]
     [:div
      [:label {:for "email", :class "sr-only"} "Email"]
      [:div {:class "relative"}
       [:input {:type "email", :class "cfi", :placeholder "Enter your email"}]
       [:span {:class "cfps"} "✉"]]]
     [:button {:type "input", :class "cfb"} "Reset password"]
     [:p {:class "cfm"} "No Account? "
      [:a {:class "underline", :href "#/register"} "Sign up"]]
     ]]])


(defn render-reset-password
  []
  (ui/render-ui reset-password-page "app"))