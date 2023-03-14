(ns aupro.profile
   (:require [centipair.ui :as ui]
             [reagent.core :as r]
             [centipair.ajax :as ajax]
             [centipair.spa :as spa]
             [centipair.components.input :as input]
             [centipair.components.notifier :as notifier]
             [aupro.form :as form]))


(def profile-id (r/atom {:id "id" :value 0}))
(def full-name (r/atom {:id "full-name" :type "text" :class "cfi" :placeholder "Enter Full Name"}))
(def about (r/atom {:id "about" :type "text-area" :class "textarea textarea-bordered textarea-lg w-full max-w-xs" :placeholder "About"}))
(def instagram (r/atom {:id "instagram" :type "text" :class "cfi" :placeholder "Instagram link"}))
(def facebook (r/atom {:id "facebook" :type "text" :class "cfi" :placeholder "Facebook link"}))
(def twitter (r/atom {:id "twitter" :type "text" :class "cfi" :placeholder "Twitter link"}))
(def youtube (r/atom {:id "youtube" :type "text" :class "cfi" :placeholder "Youtube link"}))
(def tiktok (r/atom {:id "tiktok" :type "text" :class "cfi" :placeholder "Tiktok link"}))
(def profile-pic (r/atom {:id "profile-pic" :url ""}))

(defn save-profile[])

(def save-profile-button (r/atom {:label "Save" :on-click save-profile}))
(defn profile-form 
  []
  [:div
   [:div {:class (if (> (:value @profile-id) 0)  "" "hidden")}[:img {:src "http://t3.gstatic.com/licensed-image?q=tbn:ANd9GcQcKtPg4LQ1A7_j_7_ph7FfTTTjQrnqOdC2EPUHdeqAZ01JOImw19i9gvYHROXo0HahI13E_dZ1ZekfGEE"
                :class "mx-auto shadow-xl border-solid border-2 border-gray-300 w-44 rounded-lg"}]
    (form/file profile-pic)]
   (form/generate-form "Profile"
                       "Update profile details"
                       [full-name about instagram facebook twitter youtube tiktok]
                       save-profile-button
                       [])])



(defn new-profile-form
  [] 
  (ui/render profile-form "app"))



(defn get-profile
  [id])

(defn edit-profile-form
  [id]
  (ui/render profile-form "app"))