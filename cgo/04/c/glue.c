#include <stdlib.h>
#include <stdio.h>
#include <Python.h>

void init_python() {
  Py_Initialize();
}

PyObject* import_func(char *file_path, char *func_name) {
    // PyObject *PyImport_ImportModule(const char *name)¶
    PyObject *module = PyImport_ImportModule(file_path);
    if (module == NULL) {
        return NULL;
    }

    // PyObject *PyObject_GetAttrString(PyObject *o, const char *attr_name)
    PyObject *func = PyObject_GetAttrString(module, func_name);
    if (func == NULL) {
        return NULL;
    }
    // Py_DECREF(module);
    return func;
}

PyObject *call_my_func(PyObject* func) {
    // PyObject *PyObject_CallObject(PyObject *callable, PyObject *args)
    PyObject *out = PyObject_CallObject(func, NULL);
    if (out == NULL) {
        return NULL;
    }

    return out;
}

const char *call_my_func_ret_str(char *file_path, char *func_name, char* arg) {
    PyObject *func = import_func(file_path, func_name);
    if (func == NULL) {
        printf("import_func error\n");
        return 0;
    }

    // build function args and call function
    PyObject *args = PyTuple_New(1);
    // PyObject *PyUnicode_FromStringAndSize(const char *u, Py_ssize_t size)
    PyObject * ss = PyUnicode_FromStringAndSize(arg, strlen(arg));
    PyTuple_SetItem(args, 0, ss);
    // PyObject *PyObject_CallObject(PyObject *callable, PyObject *args)
    PyObject *out = PyObject_CallObject(func, args);
    if (out == NULL) {
        return NULL;
    }

    //int PyUnicode_CheckExact(PyObject *o)¶
    if (PyUnicode_CheckExact(out) != 1) {
        printf("PyUnicode_CheckExact error\n");
        return NULL;
    }

    // PyObject *PyUnicode_AsUTF8String(PyObject *unicode)
    PyObject *some = PyUnicode_AsUTF8String(out);
    if (some == NULL) {
        printf("PyUnicode_AsUTF8String error\n");
        return NULL;
    }

    // char *PyBytes_AsString(PyObject *o)¶
    char* str = PyBytes_AsString(some);
    if (str == NULL) {
        printf("PyBytes_AsString error\n");
        return NULL;
    }

    return str;
}

const char *py_last_error() {
  PyObject *err = PyErr_Occurred();
  if (err == NULL) {
    return NULL;
  }

  PyObject *str = PyObject_Str(err);
  const char *utf8 = PyUnicode_AsUTF8(str);
  Py_DECREF(str);
  return utf8;
}

long call_my_func_ret_num(char *file_path, char *func_name) {
    PyObject *func = import_func(file_path, func_name);
    if (func == NULL) {
        printf("import_func error\n");
        return 0;
    }

    // PyObject *PyObject_CallObject(PyObject *callable, PyObject *args)
    PyObject *out = PyObject_CallObject(func, NULL);
    if (out == NULL) {
        return 0;
    }

    //int PyUnicode_CheckExact(PyObject *o)¶
    if (PyLong_Check(out) != 1) {
        printf("PyLong_Check error\n");
        return 0;
    }

    // long PyLong_AsLong(PyObject *obj)
    long num = PyLong_AsLong(out);
    if (num == -1) {
        return 0;
    }

    return num;
}

