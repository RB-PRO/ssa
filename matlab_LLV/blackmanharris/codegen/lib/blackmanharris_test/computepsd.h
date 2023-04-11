/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * computepsd.h
 *
 * Code generation for function 'computepsd'
 *
 */

#ifndef COMPUTEPSD_H
#define COMPUTEPSD_H

/* Include files */
#include <stddef.h>
#include <stdlib.h>
#include "rtwtypes.h"
#include "blackmanharris_test_types.h"

/* Function Declarations */
extern void computepsd(const double Sxx1_data[], const double w2_data[], const
  int w2_size[1], const char range[8], double varargout_1_data[], int
  varargout_1_size[1], double varargout_2_data[], int varargout_2_size[1], char
  varargout_3_data[], int varargout_3_size[2]);

#endif

/* End of code generation (computepsd.h) */
