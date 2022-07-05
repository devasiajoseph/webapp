(ns centipair.time
  (:require [cljs-time.format :as fmt]
            [cljs-time.local :as lt]))

(def indian-date-formatter (fmt/formatter "dd-MM-yyyy"))

(defn local-string-date-today
  []
  (fmt/unparse indian-date-formatter (lt/local-now)))

(defn parse-indian-date
  [date]
  (try
    (fmt/parse indian-date-formatter date)
    (catch js/Error e
      e)))
