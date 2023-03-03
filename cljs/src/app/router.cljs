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