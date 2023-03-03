(ns app.router
  (:require [reagent.core :as r] 
            [app.home :as home]
            [aupro.user :as user]
            [centipair.ui :as ui]
            [aupro.dashboard :as dash]
            [secretary.core :as secretary :refer-macros [defroute]]
            [goog.events :as events]
            [goog.history.EventType :as HistoryEventType])
  
  (:import goog.History))

(defonce match (r/atom nil))




(defn test-query 
  [qp]
  (println qp)
  )

(def routes
  [["/"
    {:name ::frontpage
     :view home/render-home
     :controllers [{:start home/render-home
                    :stop (fn [] )}]}]

   ["/about"
    {:name ::about
     :controllers [{:start home/render-about}]}]
   
   ["/login"
    {:name ::login
     :controllers [{:start user/render-login}]}]
   ["/register"
    {:name ::register
     :controllers [{:start user/render-register}]}]
   ["/reset-password"
    {:name ::reset-password
     :controllers [{:start user/render-reset-password}]}]
   ["/dashboard"
    {:name ::dashboard
     :controllers [{:start dash/render-dashboard}]}] 
   ])




(defroute home "/" [] (home/render-home))
(defroute login "/login" [] (user/render-login))
(defroute register "/register" [] (user/render-register))
(defroute reset-password "/reset-password" [] (user/render-reset-password))
(defroute dashboard "/dashboard" [] (dash/render-dashboard))
(defroute "*" [] (ui/render-ui (fn [] [:h2 "404 Not Found"]) "app"))

(defn hook-browser-navigation! []
  (doto (History.)
    (events/listen
     HistoryEventType/NAVIGATE
     (fn [event]
       (secretary/dispatch! (.-token event))))
    (.setEnabled true)))

(defn init! []
  (hook-browser-navigation!))