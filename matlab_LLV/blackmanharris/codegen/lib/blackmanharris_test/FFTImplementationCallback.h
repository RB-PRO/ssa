/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * FFTImplementationCallback.h
 *
 * Code generation for function 'FFTImplementationCallback'
 *
 */

#ifndef FFTIMPLEMENTATIONCALLBACK_H
#define FFTIMPLEMENTATIONCALLBACK_H

/* Include files */
#include <stddef.h>
#include <stdlib.h>
#include "rtwtypes.h"
#include "blackmanharris_test_types.h"

/* Function Declarations */
extern void c_FFTImplementationCallback_doH(const double x_data[], creal_T
  y_data[], const double costab_data[], const double sintab_data[]);
extern void c_FFTImplementationCallback_gen(double costab_data[], int
  costab_size[2], double sintab_data[], int sintab_size[2], int sintabinv_size[2]);

#endif

/* End of code generation (FFTImplementationCallback.h) */
