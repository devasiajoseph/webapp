;; Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
;; Unauthorized copying of this file, via any medium is strictly prohibited
;; Proprietary and confidential
;; Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018

;; ==================
;; Purpose and Design
;; ==================
;; HTML Input components
(ns centipair.components.input
  (:require [centipair.error :refer [append-field-error]]
            [centipair.components.notifier :as notifier]))


(defn update-value
  "Updates value of the input element atom"
  [field value]
  (swap! field assoc :value value)
  (if (:on-change @field)
    ((:on-change @field) field)))


(defn make-valid
  [field]
  (swap! field assoc :message "") true)

(defn valid-field?
  [field]
  (if (nil? (:validator @field))
    (do
      (make-valid field)
      true)
    (let [result ((:validator @field) @field)] 
      (if (:valid result)
        (make-valid field)
        (do
          (append-field-error field (:message result))
          false)))))



(defn valid-form?
  [form-fields]
  (apply = true (doall (map valid-field? form-fields))))


(defn perform-action
  [action form-fields]
  (notifier/notify 200)
  (if (valid-form? form-fields)
    (action)
    (notifier/notify 422 "Invalid data submitted")))

(defn text [field]
  ^{:key (:id @field)}
  [:input {:class "input input-bordered w-full"
           :type (:type @field)
           :id (:id @field) 
           :value (:value @field)
           :on-change #(update-value field (-> % .-target .-value))
           :maxLength (or (:maxLength @field) 255)
           :disabled (if (:disabled @field) "disabled" "")
           :placeholder (if (nil? (:placeholder @field))
                          ""
                          (:placeholder @field))}])


(defn text-area
  [field]
  ^{:key (:id @field)}
  [:textarea {:class (:class @field)  
              :id (:id @field)
              :value (:value @field)
              :on-change #(update-value field (-> % .-target .-value))
              :disabled (if (:disabled @field) "disabled" "")
              :placeholder (if (nil? (:placeholder @field))
                             ""
                             (:placeholder @field))}])

(defn button [action-button form-fields]
  [:div {:class "mb-3"} [:a {:class "btn btn-primary w-100"
                             :on-click #(perform-action (:on-click @action-button) form-fields)
                             :disabled (if (nil? (:disabled action-button)) "" "disabled")
                             :key (:id @action-button)}
                         (:label @action-button)]])
