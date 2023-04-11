/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * computeperiodogram.h
 *
 * Code generation for function 'computeperiodogram'
 *
 */

#ifndef COMPUTEPERIODOGRAM_H
#define COMPUTEPERIODOGRAM_H

/* Include files */
#include <stddef.h>
#include <stdlib.h>
#include "rtwtypes.h"
#include "blackmanharris_test_types.h"

/* Function Declarations */
extern void computeperiodogram(const double x_data[], const int x_size[2],
  double Pxx_data[], int Pxx_size[1], double F_data[], int F_size[1], int
  RPxx_size[1], int Fc_size[2]);

#endif

/* End of code generation (computeperiodogram.h) */
