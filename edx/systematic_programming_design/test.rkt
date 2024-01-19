#! /opt/homebrew/bin/racket
#lang racket/base

; Pythagorean Theorem of a triangle

; Side A: 3
; Side B: 4
; Side C: ?


(sqrt (+ (sqr 3) (sqr 4)))

(/ (+ 4 6.2 -12) 3 )

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; PROBLEM:
; Design a function called yell that consumes a string like 'hello'
; and adds '!' to the end of it to produce 'hello!'
;
; String -> String
; Adds "!" to the end of the input

(check-expect (yell "hello") "hello!")
(check-expect (yell "bye") "bye!")

; (define (yell s) "")   //  stub


; (define (yell s) // template
;   (... s)
;   )

(define (yell s)
  (string-append s "!")
)

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; PROBLEM:
; Design a function called area that consumes the length
; of one side of a square and produces the area of the square


; Number -> Number
; Given length, returns area of a square
;
(check-expect (area 3) 9)
(check-expect (area 3.2) (* 3.2 3.2))
; (define (area s) 0)

; (define (area s)
;   (... s))

(define (area s)
  (* s s))
