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
  (if (nil? (:class-name @field))
    true
    (do
      (swap! field assoc
             :message ""
             :class-name (clojure.string/replace (:class-name @field) #" has-error" ""))
      true)))

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
  [:input {:class (:class @field)
           :type (:type @field)
           :id (:id @field) 
           :value (:value @field)
           :on-change #(update-value field (-> % .-target .-value))
           :maxLength (or (:maxLength @field) 255)
           :disabled (if (:disabled @field) "disabled" "")
           :placeholder (if (nil? (:placeholder @field))
                          ""
                          (:placeholder @field))}])

(defn button [action-button form-fields]
  [:div {:class "mt-3"} [:a {:class "btn btn-primary"
                             :on-click #(perform-action (:on-click @action-button) form-fields)
                             :disabled (if (nil? (:disabled action-button)) "" "disabled")
                             :key (:id @action-button)}
                         (:label @action-button)]])
