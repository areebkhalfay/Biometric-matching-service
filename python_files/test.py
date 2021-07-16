import face_recognition

def test():
    # Using face_recognition api
    first_face = face_recognition.load_image_file("Face Images/1.png")
    second_face = face_recognition.load_image_file("Face Images/2.png")

    first_face_encoding = face_recognition.face_encodings(first_face)[0]
    second_face_encoding = face_recognition.face_encodings(second_face)[0]

    results = face_recognition.face_distance([first_face_encoding], second_face_encoding)
    return results

if __name__ == "__main__":
    print(1 - test()[0])