;; The first three lines of this file were inserted by DrRacket. They record metadata
;; about the language level of this file in a form that our tools can easily process.
#reader(lib "htdp-beginner-reader.ss" "lang")((modname Ex3) (read-case-sensitive #t) (teachpacks ()) (htdp-settings #(#t constructor repeating-decimal #f #t none #f () #f)))
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; PROBLEM:
; Design a function called image-area that consumes an image and
; prodices the area of that image. For the area it is sufficient
; to just multiply the image's width by its height
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require 2htdp/image)

;; Image -> Natural
;; produces image's width * heigth (area)
(check-expect (image-area (rectangle 2 3 "solid" "red")) (* 2 3))

; (define (image-area img) 0) stub
;(define (image-area img)
;  (... img)
;)

(define (image-area img)
  (* (image-width img) (image-height img))
)