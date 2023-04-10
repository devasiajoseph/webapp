(ns app.router
  (:require [goog.events :as events]
            [goog.history.EventType :as HistoryEventType]
            [centipair.ui :as ui]
            [aupro.dashboard :as dash]
            [secretary.core :as secretary :refer-macros [defroute]]
            [app.home :as home]
            [aupro.user :as user]
            [aupro.feed :as feed]
            [aupro.profile :as profile]
            [aupro.manager :as manager])
  
  (:import goog.History))




(defroute home "/" [] (home/render-home))
(defroute login "/login" [] (user/render-login))
(defroute register "/register" [] (user/render-register))
(defroute activate "/activate/:key" [key] (user/render-activate key))
(defroute reset-password "/reset-password" [] (user/render-reset-password))
(defroute post "/post/:id" [id] (feed/render-post id))
(defroute dashboard "/dashboard" [] (dash/render-dashboard))
(defroute profile-new "/profile/new" [] (profile/new-profile-form))
(defroute profile-edit "/profile/edit/:id" [id] (profile/edit-profile-form id))
(defroute profile-list "/profile/list/:page" [page] (profile/render-profile-list page))
(defroute profile-search "/profile/search/:query/:page" [query page] (profile/render-profile-search query page))
(defroute manager-list "/manager/list/:page" [page] (manager/render-manager-list page))
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