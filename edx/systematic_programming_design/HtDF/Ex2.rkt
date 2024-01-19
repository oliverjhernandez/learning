;; The first three lines of this file were inserted by DrRacket. They record metadata
;; about the language level of this file in a form that our tools can easily process.
#reader(lib "htdp-beginner-reader.ss" "lang")((modname Ex2) (read-case-sensitive #t) (teachpacks ()) (htdp-settings #(#t constructor repeating-decimal #f #t none #f () #f)))
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; PROBLEM:
; Design a function called area that consumes the length
; of one side of a square and produces the area of the square


; Number -> Number
; Given length, returns area of a square
;
(check-expect (area 3) 9)
(check-expect (area 3.2) (* 3.2 3.2))
(check-expect (area 4) 16)
; (define (area s) 0)

; (define (area s)
;   (... s))

(define (area s)
  (* s s))
