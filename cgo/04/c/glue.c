#include <stdlib.h>
#include <stdio.h>
#include <Python.h>

typedef struct CorePosition {
	long RCID;
	int Company;
	unsigned long GranularityID;
	int Seniority;
	int StartDate;
	int EndDate;
	float Weight;
	float Multiplicator;
	float Inflation;
	float EstimatedUSLogSalary;
	float FProb;
	float MProb;
	float WhiteProb;
	float BlackProb;
	float HispanicProb;
	float NativeProb;
	float ApiProb;
	float MultipleProb;
	int Region;
	int Country;
	int State;
	int Msa;
	int MappedRole;
	int Soc6dTitle;
	int HighestDegree;
	float GenderR;
	int Gender;
	float EthnicityR;
	int Ethnicity;
}  CorePosition;

void init_python() {
  Py_Initialize();
}

const char *py_last_error() {
  PyObject *err = PyErr_Occurred();
  if (err == NULL) {
    return NULL;
  }

  PyObject *str = PyObject_Str(err);
  const char *utf8 = PyUnicode_AsUTF8(str);
  Py_DECREF(str);
  PyErr_Print();
  return utf8;
}


PyObject *call_func_inner(PyObject *func, CorePosition *row) {
    PyObject *py_row = Py_BuildValue(
    "{s:l,s:i,s:k,s:i,s:i,s:i,s:f,s:f,s:f,s:f,s:f,s:f,s:f,s:f,s:f,s:f,s:f,s:f,s:i,s:i,s:i,s:i,s:i,s:i,s:i,s:f,s:i,s:f,s:i}",
    "RCID",                 row->RCID,
    "Company",              row->Company,
    "GranularityID",        row->GranularityID,
    "Seniority",            row->Seniority,
    "StartDate",            row->StartDate,
    "EndDate",              row->EndDate,
    "Weight",               row->Weight,
    "Multiplicator",        row->Multiplicator,
    "Inflation",            row->Inflation,
    "EstimatedUSLogSalary", row->EstimatedUSLogSalary,
    "FProb",                row->FProb,
    "MProb",                row->MProb,
    "WhiteProb",            row->WhiteProb,
    "BlackProb",            row->BlackProb,
    "HispanicProb",         row->HispanicProb,
    "NativeProb",           row->NativeProb,
    "ApiProb",              row->ApiProb,
    "MultipleProb",         row->MultipleProb,
    "Region",               row->Region,
    "Country",              row->Country,
    "State",                row->State,
    "Msa",                  row->Msa,
    "MappedRole",           row->MappedRole,
    "Soc6dTitle",           row->Soc6dTitle,
    "HighestDegree",        row->HighestDegree,
    "GenderR",              row->GenderR,
    "Gender",               row->Gender,
    "EthnicityR",           row->EthnicityR,
    "Ethnicity",            row->Ethnicity
    );
    if (py_row == NULL) {
        py_last_error();
        return NULL;
    }
    PyObject *args = PyTuple_New(1);
    PyTuple_SetItem(args, 0, py_row);
    PyObject *out = PyObject_CallObject(func, args);
    if (out == NULL) {
        return NULL;
    }
    return out;
}


PyThreadState * save_thread() {
    PyThreadState *threadState = PyEval_SaveThread();
    return threadState;
}

void restore_thread(PyThreadState *threadState) {
    PyEval_RestoreThread(threadState);
}

CorePosition *py_process(const char *fn, CorePosition *positions, int n_pos) {
    PyGILState_STATE gstate;
    gstate = PyGILState_Ensure();


    PyObject *pGlobal = PyDict_New();
    PyObject *pLocal = PyDict_New();
    printf("dicts: %p %p\n", pLocal, pGlobal);

    // Evaluate function
    PyObject *pValue = PyRun_String(fn, Py_file_input, pGlobal, pLocal);
    if (pValue == NULL) {
        printf("PyRun_String error\n");
        PyErr_Print();
        return NULL;
    }


    // Extract function from scope
    PyObject *key, *value;
    Py_ssize_t pos = 0;
    PyObject *func = NULL;
    while (PyDict_Next(pLocal, &pos, &key, &value)) {
        PyObject *some = PyUnicode_AsUTF8String(key);
        if (some == NULL) {
            printf("PyUnicode_AsUTF8String error\n");
            PyErr_Print();
            return NULL;
        }
        char* str = PyBytes_AsString(some);
        if (str == NULL) {
            printf("PyBytes_AsString error\n");
            PyErr_Print();
            return NULL;
        }

        printf("KEY: %s\n", str);
        printf("VAL is def: %d\n", PyFunction_Check(value));
        if (PyFunction_Check(value)) {
            func = value;
        } else {
            printf("PyFunction_Check error\n");
            PyErr_Print();
            return NULL;
        }
    }

    // Call func
    CorePosition *processed_poss = malloc(sizeof(CorePosition) * n_pos);
    if (processed_poss == NULL) {
        printf("malloc error\n");
        PyErr_Print();
        return NULL;
    }
    for (int i = 0; i < n_pos; i++) {
        PyObject *ret = call_func_inner(func, &positions[i]);
        if (ret == NULL) {
            printf("call_func_inner error\n");
            PyErr_Print();
            return NULL;
        }
        // TODO: map back also other fields + add error checks
        processed_poss[i] = (CorePosition){0};
        processed_poss[i].RCID = PyLong_AsLong(PyDict_GetItemString(ret, "RCID"));
        processed_poss[i].Weight = PyFloat_AsDouble(PyDict_GetItemString(ret, "Weight"));
    }

    PyGILState_Release(gstate);

//    Py_FinalizeEx();
    return processed_poss;
}
