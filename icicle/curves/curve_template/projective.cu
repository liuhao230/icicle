#include <cuda.h>
#include "curve_config.cuh"
#include "../../primitives/projective.cuh"

extern "C" ${CURVE_NAME_U}::projective_t random_projective_${CURVE_NAME_L}()
{
  return ${CURVE_NAME_U}::projective_t::rand_host();
}

extern "C" ${CURVE_NAME_U}::projective_t projective_zero_${CURVE_NAME_L}()
{
  return ${CURVE_NAME_U}::projective_t::zero();
}

extern "C" bool projective_is_on_curve_${CURVE_NAME_L}(${CURVE_NAME_U}::projective_t *point1)
{
  return ${CURVE_NAME_U}::projective_t::is_on_curve(*point1);
}

extern "C" ${CURVE_NAME_U}::affine_t projective_to_affine_${CURVE_NAME_L}(${CURVE_NAME_U}::projective_t *point1)
{
  return ${CURVE_NAME_U}::projective_t::to_affine(*point1);
}

extern "C" ${CURVE_NAME_U}::projective_t projective_from_affine_${CURVE_NAME_L}(${CURVE_NAME_U}::affine_t *point1)
{
  return ${CURVE_NAME_U}::projective_t::from_affine(*point1);
}

extern "C" ${CURVE_NAME_U}::scalar_field_t random_scalar_${CURVE_NAME_L}()
{
  return ${CURVE_NAME_U}::scalar_field_t::rand_host();
}

extern "C" bool eq_${CURVE_NAME_L}(${CURVE_NAME_U}::projective_t *point1, ${CURVE_NAME_U}::projective_t *point2)
{
  return (*point1 == *point2) && 
  !((point1->x == ${CURVE_NAME_U}::point_field_t::zero()) && (point1->y == ${CURVE_NAME_U}::point_field_t::zero()) && (point1->z == ${CURVE_NAME_U}::point_field_t::zero())) && 
  !((point2->x == ${CURVE_NAME_U}::point_field_t::zero()) && (point2->y == ${CURVE_NAME_U}::point_field_t::zero()) && (point2->z == ${CURVE_NAME_U}::point_field_t::zero()));
}

#if defined(G2_DEFINED)
extern "C" bool eq_g2_${CURVE_NAME_L}(${CURVE_NAME_U}::g2_projective_t *point1, ${CURVE_NAME_U}::g2_projective_t *point2)
{
  return (*point1 == *point2) && 
  !((point1->x == ${CURVE_NAME_U}::g2_point_field_t::zero()) && (point1->y == ${CURVE_NAME_U}::g2_point_field_t::zero()) && (point1->z == ${CURVE_NAME_U}::g2_point_field_t::zero())) && 
  !((point2->x == ${CURVE_NAME_U}::g2_point_field_t::zero()) && (point2->y == ${CURVE_NAME_U}::g2_point_field_t::zero()) && (point2->z == ${CURVE_NAME_U}::g2_point_field_t::zero()));
}

extern "C" ${CURVE_NAME_U}::g2_projective_t random_g2_projective_${CURVE_NAME_L}()
{
  return ${CURVE_NAME_U}::g2_projective_t::rand_host();
}

extern "C" ${CURVE_NAME_U}::g2_affine_t g2_projective_to_affine_${CURVE_NAME_L}(${CURVE_NAME_U}::g2_projective_t *point1)
{
  return ${CURVE_NAME_U}::g2_projective_t::to_affine(*point1);
}

extern "C" ${CURVE_NAME_U}::g2_projective_t g2_projective_from_affine_${CURVE_NAME_L}(${CURVE_NAME_U}::g2_affine_t *point1)
{
  return ${CURVE_NAME_U}::g2_projective_t::from_affine(*point1);
}

extern "C" bool g2_projective_is_on_curve_${CURVE_NAME_L}(${CURVE_NAME_U}::g2_projective_t *point1)
{
  return ${CURVE_NAME_U}::g2_projective_t::is_on_curve(*point1);
}

#endif
