;; Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
;; Unauthorized copying of this file, via any medium is strictly prohibited
;; Proprietary and confidential
;; Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018

;; ==================
;; Purpose and Design
;; ==================
;; This file contains Single Page Application based functions

(ns centipair.spa)

(defn get-url-hash
  []
  (.-hash js/location))

(defn home-page?
  "Checks whether the current page is home page
  if home page load feault componenets"
  []
  (if (or (= "" (.-hash js/location)) (nil? (.-hash js/location)))
    true false))



(defn redirect [hash-url]
  (aset  js/window.location "hash" (str hash-url)))

(defn redirect-window [url]
  (aset js/window "location" url))


(defn do-nothing [])

(defn render-initial
  []
  (let [initial-url (get-url-hash)]
    (if (not (or (= "#/login" initial-url) (= "#/register" initial-url)))
      (redirect "/login"))))
