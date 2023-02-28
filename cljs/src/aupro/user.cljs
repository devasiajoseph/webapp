(ns aupro.user
  (:require [centipair.ui :as ui]))


(defn login-page 
  []
  [:div {:class "mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8"}
   [:div {:class "mx-auto max-w-lg bg-white pt-10 rounded-md"}
    [:h1 {:class "text-center text-2xl font-bold text-indigo-600 sm:text-3xl"} "Login"]
    [:form { :class "mt-6 mb-0 space-y-4 rounded-lg p-4 shadow-lg sm:p-6 lg:p-8"}
     [:p {:class "text-center text-lg font-medium"} "Sign in to your account"]
     [:div
      [:label {:for "email", :class "sr-only"} "Email"]
      [:div {:class "relative"}
       [:input {:type "email", :class "w-full rounded-lg border-gray-200 p-4 pr-12 text-sm shadow-sm", :placeholder "Enter email"}]
       [:span {:class "absolute inset-y-0 right-0 grid place-content-center px-4"}
       ]]]
     [:div
      [:label {:for "password", :class "sr-only"} "Password"]
      [:div {:class "relative"}
       [:input {:type "password", :class "w-full rounded-lg border-gray-200 p-4 pr-12 text-sm shadow-sm", :placeholder "Enter password"}]
       [:span {:class "absolute inset-y-0 right-0 grid place-content-center px-4"}]]]
     [:button {:type "input", :class "block w-full rounded-lg bg-indigo-600 px-5 py-3 text-sm font-medium text-white"} "Sign in"]
     [:p {:class "text-center text-sm text-gray-500"} "No account?"
      [:a {:class "underline", :href ""} "Sign up"]]]]])


(defn render-login
  []
  (ui/render-ui login-page "app"))



(defn register-page
  []
  [:div {:class "mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8"}
   [:div {:class "mx-auto max-w-lg bg-white pt-10 rounded-md"}
    [:h1 {:class "text-center text-2xl font-bold text-indigo-600 sm:text-3xl"} "Sign up"]
    [:form {:class "mt-6 mb-0 space-y-4 rounded-lg p-4 shadow-lg sm:p-6 lg:p-8"}
     [:p {:class "text-center text-lg font-medium"} "Register a new account"]
     [:div
      [:label {:for "email", :class "sr-only"} "Email"]
      [:div {:class "relative"}
       [:input {:type "email", :class "w-full rounded-lg border-gray-200 p-4 pr-12 text-sm shadow-sm", :placeholder "Enter email"}]
       [:span {:class "absolute inset-y-0 right-0 grid place-content-center px-4"}]]]
     [:div
      [:label {:for "password", :class "sr-only"} "Password"]
      [:div {:class "relative"}
       [:input {:type "password", :class "w-full rounded-lg border-gray-200 p-4 pr-12 text-sm shadow-sm", :placeholder "Enter password"}]
       [:span {:class "absolute inset-y-0 right-0 grid place-content-center px-4"}]]]
     [:button {:type "input", :class "block w-full rounded-lg bg-indigo-600 px-5 py-3 text-sm font-medium text-white"} "Register"]
     [:p {:class "text-center text-sm text-gray-500"} "No account?"
      [:a {:class "underline", :href ""} "Sign up"]]]]])


(defn render-register
  []
  (ui/render-ui register-page "app"))