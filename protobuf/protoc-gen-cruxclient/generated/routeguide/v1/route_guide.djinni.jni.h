// This file is generated by s12-proto api generator. Please DO NOT modify.
#pragma once
#include <string>
#include "djinni_support.hpp"
#include "routeguide/v1/message.pb.h"
namespace djinni::routeguide::v1::Point {
struct JNIInfo {
  const GlobalRef<jclass> clazz { jniFindClass("io/grpc/examples/routeguide/Point")  };
  const jmethodID method_toBytes { jniGetMethodID(clazz.get(), "toByteArray", "()[B") };
  const jmethodID method_byteSize { jniGetMethodID(clazz.get(), "getSerializedSize", "()I") };
  const jmethodID method_fromBytes { jniGetStaticMethodID(clazz.get(), "parseFrom", "([B)Lio/grpc/examples/routeguide/Point;") };
};

struct Translator {
  using CppType = ::routeguide::v1::Point;
  using JniType = jobject;
  using Boxed = Translator;

  static CppType toCpp(JNIEnv* jniEnv, JniType j) {
    assert(j != nullptr);
    const auto& data = JniClass<JNIInfo>::get();
    assert(jniEnv->IsInstanceOf(j, data.clazz.get()));
    jbyte b = jniEnv->CallByteMethod(j, data.method_toBytes);
    auto byte_len = jniEnv->CallIntMethod(j, data.method_byteSize);
    jniExceptionCheck(jniEnv);
    CppType cpp_message;
    cpp_message.ParseFromArray(&b, byte_len);
    return cpp_message;
  }

  static LocalRef<JniType> fromCpp(JNIEnv* jniEnv, const CppType& message) {
    size_t size = message.ByteSizeLong();
    jbyte* temp = new jbyte[size];
    message.SerializeToArray(temp, static_cast<int>(size));
    jbyteArray bytes = jniEnv->NewByteArray(size);
    jniEnv->SetByteArrayRegion(bytes, 0, size, temp);
    delete[] temp;
    const auto& data = JniClass<JNIInfo>::get();
    auto j = LocalRef<JniType>{jniEnv->CallStaticObjectMethod(data.clazz.get(), data.method_fromBytes, bytes)};
    jniExceptionCheck(jniEnv);
    return j;
  }
};
}  //namespace djinni::routeguide::v1::Point

#include "routeguide/v1/message.pb.h"
namespace djinni::routeguide::v1::Rectangle {
struct JNIInfo {
  const GlobalRef<jclass> clazz { jniFindClass("io/grpc/examples/routeguide/Rectangle")  };
  const jmethodID method_toBytes { jniGetMethodID(clazz.get(), "toByteArray", "()[B") };
  const jmethodID method_byteSize { jniGetMethodID(clazz.get(), "getSerializedSize", "()I") };
  const jmethodID method_fromBytes { jniGetStaticMethodID(clazz.get(), "parseFrom", "([B)Lio/grpc/examples/routeguide/Rectangle;") };
};

struct Translator {
  using CppType = ::routeguide::v1::Rectangle;
  using JniType = jobject;
  using Boxed = Translator;

  static CppType toCpp(JNIEnv* jniEnv, JniType j) {
    assert(j != nullptr);
    const auto& data = JniClass<JNIInfo>::get();
    assert(jniEnv->IsInstanceOf(j, data.clazz.get()));
    jbyte b = jniEnv->CallByteMethod(j, data.method_toBytes);
    auto byte_len = jniEnv->CallIntMethod(j, data.method_byteSize);
    jniExceptionCheck(jniEnv);
    CppType cpp_message;
    cpp_message.ParseFromArray(&b, byte_len);
    return cpp_message;
  }

  static LocalRef<JniType> fromCpp(JNIEnv* jniEnv, const CppType& message) {
    size_t size = message.ByteSizeLong();
    jbyte* temp = new jbyte[size];
    message.SerializeToArray(temp, static_cast<int>(size));
    jbyteArray bytes = jniEnv->NewByteArray(size);
    jniEnv->SetByteArrayRegion(bytes, 0, size, temp);
    delete[] temp;
    const auto& data = JniClass<JNIInfo>::get();
    auto j = LocalRef<JniType>{jniEnv->CallStaticObjectMethod(data.clazz.get(), data.method_fromBytes, bytes)};
    jniExceptionCheck(jniEnv);
    return j;
  }
};
}  //namespace djinni::routeguide::v1::Rectangle

#include "routeguide/v1/message.pb.h"
namespace djinni::routeguide::v1::Feature {
struct JNIInfo {
  const GlobalRef<jclass> clazz { jniFindClass("io/grpc/examples/routeguide/Feature")  };
  const jmethodID method_toBytes { jniGetMethodID(clazz.get(), "toByteArray", "()[B") };
  const jmethodID method_byteSize { jniGetMethodID(clazz.get(), "getSerializedSize", "()I") };
  const jmethodID method_fromBytes { jniGetStaticMethodID(clazz.get(), "parseFrom", "([B)Lio/grpc/examples/routeguide/Feature;") };
};

struct Translator {
  using CppType = ::routeguide::v1::Feature;
  using JniType = jobject;
  using Boxed = Translator;

  static CppType toCpp(JNIEnv* jniEnv, JniType j) {
    assert(j != nullptr);
    const auto& data = JniClass<JNIInfo>::get();
    assert(jniEnv->IsInstanceOf(j, data.clazz.get()));
    jbyte b = jniEnv->CallByteMethod(j, data.method_toBytes);
    auto byte_len = jniEnv->CallIntMethod(j, data.method_byteSize);
    jniExceptionCheck(jniEnv);
    CppType cpp_message;
    cpp_message.ParseFromArray(&b, byte_len);
    return cpp_message;
  }

  static LocalRef<JniType> fromCpp(JNIEnv* jniEnv, const CppType& message) {
    size_t size = message.ByteSizeLong();
    jbyte* temp = new jbyte[size];
    message.SerializeToArray(temp, static_cast<int>(size));
    jbyteArray bytes = jniEnv->NewByteArray(size);
    jniEnv->SetByteArrayRegion(bytes, 0, size, temp);
    delete[] temp;
    const auto& data = JniClass<JNIInfo>::get();
    auto j = LocalRef<JniType>{jniEnv->CallStaticObjectMethod(data.clazz.get(), data.method_fromBytes, bytes)};
    jniExceptionCheck(jniEnv);
    return j;
  }
};
}  //namespace djinni::routeguide::v1::Feature

#include "routeguide/v1/message.pb.h"
namespace djinni::routeguide::v1::RouteNote {
struct JNIInfo {
  const GlobalRef<jclass> clazz { jniFindClass("io/grpc/examples/routeguide/RouteNote")  };
  const jmethodID method_toBytes { jniGetMethodID(clazz.get(), "toByteArray", "()[B") };
  const jmethodID method_byteSize { jniGetMethodID(clazz.get(), "getSerializedSize", "()I") };
  const jmethodID method_fromBytes { jniGetStaticMethodID(clazz.get(), "parseFrom", "([B)Lio/grpc/examples/routeguide/RouteNote;") };
};

struct Translator {
  using CppType = ::routeguide::v1::RouteNote;
  using JniType = jobject;
  using Boxed = Translator;

  static CppType toCpp(JNIEnv* jniEnv, JniType j) {
    assert(j != nullptr);
    const auto& data = JniClass<JNIInfo>::get();
    assert(jniEnv->IsInstanceOf(j, data.clazz.get()));
    jbyte b = jniEnv->CallByteMethod(j, data.method_toBytes);
    auto byte_len = jniEnv->CallIntMethod(j, data.method_byteSize);
    jniExceptionCheck(jniEnv);
    CppType cpp_message;
    cpp_message.ParseFromArray(&b, byte_len);
    return cpp_message;
  }

  static LocalRef<JniType> fromCpp(JNIEnv* jniEnv, const CppType& message) {
    size_t size = message.ByteSizeLong();
    jbyte* temp = new jbyte[size];
    message.SerializeToArray(temp, static_cast<int>(size));
    jbyteArray bytes = jniEnv->NewByteArray(size);
    jniEnv->SetByteArrayRegion(bytes, 0, size, temp);
    delete[] temp;
    const auto& data = JniClass<JNIInfo>::get();
    auto j = LocalRef<JniType>{jniEnv->CallStaticObjectMethod(data.clazz.get(), data.method_fromBytes, bytes)};
    jniExceptionCheck(jniEnv);
    return j;
  }
};
}  //namespace djinni::routeguide::v1::RouteNote

#include "routeguide/v1/message.pb.h"
namespace djinni::routeguide::v1::RouteSummary {
struct JNIInfo {
  const GlobalRef<jclass> clazz { jniFindClass("io/grpc/examples/routeguide/RouteSummary")  };
  const jmethodID method_toBytes { jniGetMethodID(clazz.get(), "toByteArray", "()[B") };
  const jmethodID method_byteSize { jniGetMethodID(clazz.get(), "getSerializedSize", "()I") };
  const jmethodID method_fromBytes { jniGetStaticMethodID(clazz.get(), "parseFrom", "([B)Lio/grpc/examples/routeguide/RouteSummary;") };
};

struct Translator {
  using CppType = ::routeguide::v1::RouteSummary;
  using JniType = jobject;
  using Boxed = Translator;

  static CppType toCpp(JNIEnv* jniEnv, JniType j) {
    assert(j != nullptr);
    const auto& data = JniClass<JNIInfo>::get();
    assert(jniEnv->IsInstanceOf(j, data.clazz.get()));
    jbyte b = jniEnv->CallByteMethod(j, data.method_toBytes);
    auto byte_len = jniEnv->CallIntMethod(j, data.method_byteSize);
    jniExceptionCheck(jniEnv);
    CppType cpp_message;
    cpp_message.ParseFromArray(&b, byte_len);
    return cpp_message;
  }

  static LocalRef<JniType> fromCpp(JNIEnv* jniEnv, const CppType& message) {
    size_t size = message.ByteSizeLong();
    jbyte* temp = new jbyte[size];
    message.SerializeToArray(temp, static_cast<int>(size));
    jbyteArray bytes = jniEnv->NewByteArray(size);
    jniEnv->SetByteArrayRegion(bytes, 0, size, temp);
    delete[] temp;
    const auto& data = JniClass<JNIInfo>::get();
    auto j = LocalRef<JniType>{jniEnv->CallStaticObjectMethod(data.clazz.get(), data.method_fromBytes, bytes)};
    jniExceptionCheck(jniEnv);
    return j;
  }
};
}  //namespace djinni::routeguide::v1::RouteSummary

