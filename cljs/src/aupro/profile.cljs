(ns aupro.profile
   (:require [centipair.ui :as ui]
             [reagent.core :as r]
             [centipair.ajax :as ajax]
             [centipair.spa :as spa]
             [centipair.components.input :as input]
             [centipair.components.notifier :as notifier]))



(def full-name (r/atom {:id "full-name" :type "text" :class "cfi" :placeholder "Enter Full Name"}))
(def about (r/atom {:id "about" :type "text-area" :class "cfi"}))
(def instagram (r/atom {:id "instagram" :type "text" :class "cfi" :placeholder "Instagram link"}))
(def facebook (r/atom {:id "facebook" :type "text" :class "cfi" :placeholder "Facebook link"}))
(def twitter (r/atom {:id "twitter" :type "text" :class "cfi" :placeholder "Twitter link"}))
(def youtube (r/atom {:id "youtube" :type "text" :class "cfi" :placeholder "Youtube link"}))
(def tiktok (r/atom {:id "tiktok" :type "text" :class "cfi" :placeholder "Tiktok link"}))


(def save-profile[])

(def save-profile-button (r/atom {:label "Login" :on-click save-profile}))
(defn profile-form 
  []
)



(defn render-profile-form
  [])