(ns aupro.manager
  (:require  [centipair.ui :as ui]
             [reagent.core :as r]
             [centipair.ajax :as ajax]
             [centipair.components.input :as input]
             [centipair.components.search :as search]
             [centipair.components.pagination :as p]))


(def manager-list (r/atom {:url "#/manager/list/" :limit 50}))
(def manager (r/atom {:id "manager" :type "text" :class "cfi" :placeholder "Enter email or phone" :label "Email or Phone"}))
(def manager-search (r/atom {}))


(defn search[])

(defn manager-table
  []
  [:div {:class "overflow-x-auto w-full mb-10"}
   [:div {:class "max-w-2xl mx-auto"}
    (search/search-box manager-search search)
    [:div (p/view manager-list) [:label {:class "btn btn-primary" :for "modal-fn"} "New +"]]
    [:table {:class "table w-full"}
     [:thead [:tr [:th "Name"] [:th "Email"] [:th "Phone"]]]
     [:tbody
      (doall (map (fn [each]  ^{:key each}
                    [:tr
                     [:td [:div {:class "flex items-center space-x-3"}
                           [:div {:class "avatar"}
                            [:div {:class "mask mask-squircle w-12 h-12"}
                             [:img {:src (:profile_pic each), :alt "Avatar Tailwind CSS Component"}]]]]]
                     [:td [:a {:href (str "#/profile/edit/" (:profile_id each)) :class "btn btn-ghost btn-xs"} (:full_name each)]]]) (:data @manager-list)))]]
    (p/view manager-list)]])


(defn list-managers
  []
  )


(defn render-manager-list
  [page]
  (ui/render manager-table "app"))

(defn add-manager
  []
  (ajax/form-post "/api/manager" [manager] (fn [response])))