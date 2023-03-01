(ns app.router
  (:require [reagent.core :as r]
            [reitit.frontend :as rf]
            [reitit.frontend.easy :as rfe]
            [reitit.frontend.controllers :as rfc]
            [reitit.coercion.schema :as rsc]
            [app.home :as home]
            [aupro.user :as user]
            [aupro.dashboard :as dash]
            ))

(defonce match (r/atom nil))


(def routes
  [["/"
    {:name ::frontpage 
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

(defn init! []
  (rfe/start!
   (rf/router routes {:data {:coercion rsc/coercion}})
   (fn [new-match]
     (swap! match (fn [old-match]
                    (if new-match
                      (assoc new-match :controllers (rfc/apply-controllers (:controllers old-match) new-match))))))
    ;; set to false to enable HistoryAPI
   {:use-fragment true})
  ;;(rdom/render [home-page] (.getElementById js/document "app"))
  )