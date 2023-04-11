/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * blackmanharris_test_initialize.c
 *
 * Code generation for function 'blackmanharris_test_initialize'
 *
 */

/* Include files */
#include "blackmanharris_test_initialize.h"
#include "blackmanharris_test.h"
#include "blackmanharris_test_data.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
void blackmanharris_test_initialize(void)
{
  rt_InitInfAndNaN();
  isInitialized_blackmanharris_test = true;
}

/* End of code generation (blackmanharris_test_initialize.c) */
