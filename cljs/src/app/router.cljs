(ns app.router
  (:require [centipair.ajax :as ajax]
            [app.home :as home]
            [aupro.user :as user]
            [aupro.feed :as feed]
            [centipair.ui :as ui]
            [aupro.dashboard :as dash]
            [secretary.core :as secretary :refer-macros [defroute]]
            [goog.events :as events]
            [goog.history.EventType :as HistoryEventType])
  
  (:import goog.History))




(defroute home "/" [] (home/render-home))
(defroute login "/login" [] (user/render-login))
(defroute register "/register" [] (user/render-register))
(defroute activate "/activate/:key" [key] (user/render-activate key))
(defroute reset-password "/reset-password" [] (user/render-reset-password))
(defroute post "/post/:id" [id] (feed/render-post id))
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
  (user/fetch-menu)
  (hook-browser-navigation!))