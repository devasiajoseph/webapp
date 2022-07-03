(ns app.home
   (:require [reagent.dom :as rdom]
             [reagent.core :as r]
             [reitit.frontend :as rf]
             [reitit.frontend.easy :as rfe]
             [reitit.frontend.controllers :as rfc]
             [reitit.coercion.schema :as rsc]
             ))

(defonce match (r/atom nil))

(defn home-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to APP"]
   [:p "This page is rendered from clojurescript reagent"]])


(defn render-home []
  (print "*******" @match)
  (rdom/render [home-page] (.getElementById js/document "app")))

(defn about-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to about"]
   [:p "This is about page"]])


(defn render-about []
  (print "*******" @match)
  (rdom/render [about-page] (.getElementById js/document "app")))


(defn log-fn [& params]
  (fn [_]
    (apply js/console.log params)))

(def routes
  [["/"
    {:name ::frontpage
     :view home-page
     :controllers [{:start render-home
                    :stop (fn [] (print "stopped home"))}]}]

   ["/about"
    {:name ::about
     :view about-page
     :controllers [{:start render-about}]}]])

(defn current-page []
  [:div
   [:ul
    [:li [:a {:href (rfe/href ::frontpage)} "Frontpage"]]
    [:li
     [:a {:href (rfe/href ::item-list)} "Item list"]]]
   (if @match
     (let [view (:view (:data @match))]
       [view @match]))
   [:pre (with-out-str (print @match))]])

(defn init! []
  (rfe/start!
   (rf/router routes {:data {:coercion rsc/coercion}})
   (fn [new-match]
    (print (:query-params new-match) ) 
     (swap! match (fn [old-match]
                    (print "+++++++" old-match)
                    (if new-match
                      (assoc new-match :controllers (rfc/apply-controllers (:controllers old-match) new-match))))))
    ;; set to false to enable HistoryAPI
   {:use-fragment true})
  ;;(rdom/render [home-page] (.getElementById js/document "app"))
  )