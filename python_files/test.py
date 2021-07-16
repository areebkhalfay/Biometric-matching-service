import face_recognition
import cv2
import numpy as np
from face_compare.images import get_face
from face_compare.model import facenet_model, img_to_encoding
import tensorflow as tf

path1 = "/home/areebk/go/src/SAICCodingAssessment/Face Images/1.png"
path2 = "/home/areebk/go/src/SAICCodingAssessment/Face Images/2.png"
path3 = "/home/areebk/go/src/SAICCodingAssessment/Face Images/3.png"
path4 = "/home/areebk/go/src/SAICCodingAssessment/Face Images/4.png"
path5 = "/home/areebk/go/src/SAICCodingAssessment/Face Images/5.png"
path6 = "/home/areebk/go/src/SAICCodingAssessment/Face Images/6.png"


def test():
    # Using face_recognition api
    first_face = face_recognition.load_image_file(path1)
    second_face = face_recognition.load_image_file(path2)

    first_face_encoding = face_recognition.face_encodings(first_face)[0]
    second_face_encoding = face_recognition.face_encodings(second_face)[0]

    results = face_recognition.face_distance([first_face_encoding], second_face_encoding)
    return 1 - results


def test1():
    # Using OpenCv and face_compare apis
    # Repurposed code from face_compare api, compare_faces.py file
    face_one = get_face(cv2.imread(path3, 1))
    face_two = get_face(cv2.imread(path4, 1))

    model = facenet_model(input_shape=(3, 96, 96))

    embedding_one = img_to_encoding(face_one, model)
    embedding_two = img_to_encoding(face_two, model)

    dist = np.linalg.norm(embedding_one - embedding_two)

    return dist


if __name__ == "__main__":
    print(test()[0])
    print(test1())
