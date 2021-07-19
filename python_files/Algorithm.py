import string

import face_recognition
import base64
from __future__ import annotations


def match_images(string1: string, string2: string) -> float:
    # Using face_recognition api

    # Convert Strings to images
    image1 = base64.b64decode(string1)
    # Would need to change to directory on home machine
    with open('/home/areebk/go/src/SAICCodingAssessment/Test Images/1.png', 'wb') as f:
        f.write(image1)
    image2 = base64.b64decode(string2)
    # Would need to change to directory on home machine
    with open('/home/areebk/go/src/SAICCodingAssessment/Test Images/2.png', 'wb') as f:
        f.write(image2)
    # Would need to change to directory on home machine
    first_face = face_recognition.load_image_file("/home/areebk/go/src/SAICCodingAssessment/Test Images/1.png")Fixd
    # Would need to change to directory on home machine
    second_face = face_recognition.load_image_file("/home/areebk/go/src/SAICCodingAssessment/Test Images/2.png")

    first_face_encoding = face_recognition.face_encodings(first_face)[0]
    second_face_encoding = face_recognition.face_encodings(second_face)[0]

    results = face_recognition.face_distance([first_face_encoding], second_face_encoding)
    return (1 - results)[0]
