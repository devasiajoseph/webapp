(ns centipair.core
  (:require [app.router :as app]
            [centipair.components.notifier :as notifier]))


;;(app/render-home)
(app/init!)
(notifier/render-notifier-component)