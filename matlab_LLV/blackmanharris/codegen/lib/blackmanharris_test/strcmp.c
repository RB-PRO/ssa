/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * strcmp.c
 *
 * Code generation for function 'strcmp'
 *
 */

/* Include files */
#include "strcmp.h"
#include "blackmanharris_test.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
boolean_T b_strcmp(const char a[8])
{
  int ret;
  static const char b[8] = { 'o', 'n', 'e', 's', 'i', 'd', 'e', 'd' };

  ret = memcmp(&a[0], &b[0], 8);
  return ret == 0;
}

/* End of code generation (strcmp.c) */
