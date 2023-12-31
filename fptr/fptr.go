/**
 * AUTO GENERATED FILE
 */

package fptr10

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"time"
	"unsafe"
)

/*
#include "libfptr10_go.h"

int bridge_libfptr_create_func(libfptr_create_func f, libfptr_handle *fptr) {
    return f(fptr);
}

int bridge_libfptr_create_with_id_func(libfptr_create_with_id_func f, libfptr_handle *fptr, const wchar_t *id) {
    return f(fptr, id);
}

void bridge_libfptr_destroy_func(libfptr_destroy_func f, libfptr_handle *fptr) {
    f(fptr);
}

const char* bridge_libfptr_get_version_string_func(libfptr_get_version_string_func f, libfptr_handle fptr) {
    return f(fptr);
}

int bridge_libfptr_set_settings_func(libfptr_set_settings_func f, libfptr_handle fptr, const wchar_t *settings) {
    return f(fptr, settings);
}

int bridge_libfptr_get_settings_func(libfptr_get_settings_func f, libfptr_handle fptr, const wchar_t *value, int size) {
    return f(fptr, value, size);
}

void bridge_libfptr_set_single_setting_func(libfptr_set_single_setting_func f, libfptr_handle fptr, const wchar_t *key, const wchar_t *value) {
    return f(fptr, key, value);
}

int bridge_libfptr_get_single_setting_func(libfptr_get_single_setting_func f, libfptr_handle fptr, const wchar_t *key, wchar_t *value, int size) {
    return f(fptr, key, value, size);
}

int bridge_libfptr_is_opened_func(libfptr_is_opened_func f, libfptr_handle fptr) {
    return f(fptr);
}

int bridge_libfptr_error_code_func(libfptr_error_code_func f, libfptr_handle fptr) {
    return f(fptr);
}

int bridge_libfptr_error_description_func(libfptr_error_description_func f, libfptr_handle fptr, wchar_t *value, int size) {
    return f(fptr, value, size);
}

int bridge_libfptr_error_recommendation_func(libfptr_error_recommendation_func f, libfptr_handle fptr, wchar_t *value, int size) {
    return f(fptr, value, size);
}

int bridge_libfptr_reset_error_func(libfptr_reset_error_func f, libfptr_handle fptr) {
    return f(fptr);
}

void bridge_libfptr_set_param_bool_func(libfptr_set_param_bool_func f, libfptr_handle fptr, int param_id, int value) {
    return f(fptr, param_id, value);
}

void bridge_libfptr_set_param_int_func(libfptr_set_param_int_func f, libfptr_handle fptr, int param_id, uint value) {
    return f(fptr, param_id, value);
}

void bridge_libfptr_set_param_double_func(libfptr_set_param_double_func f, libfptr_handle fptr, int param_id, double value) {
    return f(fptr, param_id, value);
}

void bridge_libfptr_set_param_str_func(libfptr_set_param_str_func f, libfptr_handle fptr, int param_id, const wchar_t *value) {
    return f(fptr, param_id, value);
}

void bridge_libfptr_set_param_bytearray_func(libfptr_set_param_bytearray_func f, libfptr_handle fptr, int param_id, const unsigned char *value, int size) {
    return f(fptr, param_id, value, size);
}

void bridge_libfptr_set_param_datetime_func(libfptr_set_param_datetime_func f, libfptr_handle fptr, int param_id, int year, int month, int day, int hour, int minute, int second) {
    return f(fptr, param_id, year, month, day, hour, minute, second);
}

int bridge_libfptr_get_param_bool_func(libfptr_get_param_bool_func f, libfptr_handle fptr, int param_id) {
    return f(fptr, param_id);
}

uint bridge_libfptr_get_param_int_func(libfptr_get_param_int_func f, libfptr_handle fptr, int param_id) {
    return f(fptr, param_id);
}

double bridge_libfptr_get_param_double_func(libfptr_get_param_double_func f, libfptr_handle fptr, int param_id) {
    return f(fptr, param_id);
}

int bridge_libfptr_get_param_str_func(libfptr_get_param_str_func f, libfptr_handle fptr, int param_id, wchar_t *value, int size) {
    return f(fptr, param_id, value, size);
}

int bridge_libfptr_get_param_bytearray_func(libfptr_get_param_bytearray_func f, libfptr_handle fptr, int param_id, unsigned char *value, int size) {
    return f(fptr, param_id, value, size);
}

void bridge_libfptr_get_param_datetime_func(libfptr_get_param_datetime_func f, libfptr_handle fptr, int param_id, int *year, int *month, int *day, int *hour, int *minute, int *second) {
    return f(fptr, param_id, year, month, day, hour, minute, second);
}

int bridge_libfptr_is_param_available_func(libfptr_is_param_available_func f, libfptr_handle fptr, int param_id) {
    return f(fptr, param_id);
}

int bridge_libfptr_log_write_func(libfptr_log_write_func f, libfptr_handle fptr, const wchar_t *tag, int level, const wchar_t *message) {
    return f(fptr, tag, level, message);
}

int bridge_libfptr_change_label_func(libfptr_change_label_func f, libfptr_handle fptr, const wchar_t *label) {
    return f(fptr, label);
}

int bridge_libfptr_show_properties_func(libfptr_show_properties_func f, libfptr_handle fptr, uint parent_type, void *parent) {
    return f(fptr, parent_type, parent);
}

int bridge_libfptr_simple_func(libfptr_simple_call_func f, libfptr_handle fptr) {
    return f(fptr);
}
*/
import "C"

const (
	LIBFPTR_PARAM_TEXT                                             = 65536
	LIBFPTR_PARAM_TEXT_WRAP                                        = 65537
	LIBFPTR_PARAM_ALIGNMENT                                        = 65538
	LIBFPTR_PARAM_FONT                                             = 65539
	LIBFPTR_PARAM_FONT_DOUBLE_WIDTH                                = 65540
	LIBFPTR_PARAM_FONT_DOUBLE_HEIGHT                               = 65541
	LIBFPTR_PARAM_LINESPACING                                      = 65542
	LIBFPTR_PARAM_BRIGHTNESS                                       = 65543
	LIBFPTR_PARAM_MODEL                                            = 65544
	LIBFPTR_PARAM_RECEIPT_TYPE                                     = 65545
	LIBFPTR_PARAM_REPORT_TYPE                                      = 65546
	LIBFPTR_PARAM_MODE                                             = 65547
	LIBFPTR_PARAM_EXTERNAL_DEVICE_TYPE                             = 65548
	LIBFPTR_PARAM_EXTERNAL_DEVICE_DATA                             = 65549
	LIBFPTR_PARAM_FREQUENCY                                        = 65550
	LIBFPTR_PARAM_DURATION                                         = 65551
	LIBFPTR_PARAM_CUT_TYPE                                         = 65552
	LIBFPTR_PARAM_DRAWER_ON_TIMEOUT                                = 65553
	LIBFPTR_PARAM_DRAWER_OFF_TIMEOUT                               = 65554
	LIBFPTR_PARAM_DRAWER_ON_QUANTITY                               = 65555
	LIBFPTR_PARAM_TIMEOUT_ENQ                                      = 65556
	LIBFPTR_PARAM_COMMAND_BUFFER                                   = 65557
	LIBFPTR_PARAM_ANSWER_BUFFER                                    = 65558
	LIBFPTR_PARAM_SERIAL_NUMBER                                    = 65559
	LIBFPTR_PARAM_MANUFACTURER_CODE                                = 65560
	LIBFPTR_PARAM_NO_NEED_ANSWER                                   = 65561
	LIBFPTR_PARAM_INFO_DISCOUNT_SUM                                = 65562
	LIBFPTR_PARAM_USE_ONLY_TAX_TYPE                                = 65563
	LIBFPTR_PARAM_PAYMENT_TYPE                                     = 65564
	LIBFPTR_PARAM_PAYMENT_SUM                                      = 65565
	LIBFPTR_PARAM_REMAINDER                                        = 65566
	LIBFPTR_PARAM_CHANGE                                           = 65567
	LIBFPTR_PARAM_DEPARTMENT                                       = 65568
	LIBFPTR_PARAM_TAX_TYPE                                         = 65569
	LIBFPTR_PARAM_TAX_SUM                                          = 65570
	LIBFPTR_PARAM_TAX_MODE                                         = 65571
	LIBFPTR_PARAM_RECEIPT_ELECTRONICALLY                           = 65572
	LIBFPTR_PARAM_USER_PASSWORD                                    = 65573
	LIBFPTR_PARAM_SCALE                                            = 65574
	LIBFPTR_PARAM_LEFT_MARGIN                                      = 65575
	LIBFPTR_PARAM_BARCODE                                          = 65576
	LIBFPTR_PARAM_BARCODE_TYPE                                     = 65577
	LIBFPTR_PARAM_BARCODE_PRINT_TEXT                               = 65578
	LIBFPTR_PARAM_BARCODE_VERSION                                  = 65579
	LIBFPTR_PARAM_BARCODE_CORRECTION                               = 65580
	LIBFPTR_PARAM_BARCODE_COLUMNS                                  = 65581
	LIBFPTR_PARAM_BARCODE_INVERT                                   = 65582
	LIBFPTR_PARAM_HEIGHT                                           = 65583
	LIBFPTR_PARAM_WIDTH                                            = 65584
	LIBFPTR_PARAM_FILENAME                                         = 65585
	LIBFPTR_PARAM_PICTURE_NUMBER                                   = 65586
	LIBFPTR_PARAM_DATA_TYPE                                        = 65587
	LIBFPTR_PARAM_OPERATOR_ID                                      = 65588
	LIBFPTR_PARAM_LOGICAL_NUMBER                                   = 65589
	LIBFPTR_PARAM_DATE_TIME                                        = 65590
	LIBFPTR_PARAM_FISCAL                                           = 65591
	LIBFPTR_PARAM_SHIFT_STATE                                      = 65592
	LIBFPTR_PARAM_CASHDRAWER_OPENED                                = 65593
	LIBFPTR_PARAM_RECEIPT_PAPER_PRESENT                            = 65594
	LIBFPTR_PARAM_COVER_OPENED                                     = 65595
	LIBFPTR_PARAM_SUBMODE                                          = 65596
	LIBFPTR_PARAM_RECEIPT_NUMBER                                   = 65597
	LIBFPTR_PARAM_DOCUMENT_NUMBER                                  = 65598
	LIBFPTR_PARAM_SHIFT_NUMBER                                     = 65599
	LIBFPTR_PARAM_RECEIPT_SUM                                      = 65600
	LIBFPTR_PARAM_RECEIPT_LINE_LENGTH                              = 65601
	LIBFPTR_PARAM_RECEIPT_LINE_LENGTH_PIX                          = 65602
	LIBFPTR_PARAM_MODEL_NAME                                       = 65603
	LIBFPTR_PARAM_UNIT_VERSION                                     = 65604
	LIBFPTR_PARAM_PRINTER_CONNECTION_LOST                          = 65605
	LIBFPTR_PARAM_PRINTER_ERROR                                    = 65606
	LIBFPTR_PARAM_CUT_ERROR                                        = 65607
	LIBFPTR_PARAM_PRINTER_OVERHEAT                                 = 65608
	LIBFPTR_PARAM_UNIT_TYPE                                        = 65609
	LIBFPTR_PARAM_LICENSE_NUMBER                                   = 65610
	LIBFPTR_PARAM_LICENSE_ENTERED                                  = 65611
	LIBFPTR_PARAM_LICENSE                                          = 65612
	LIBFPTR_PARAM_SUM                                              = 65613
	LIBFPTR_PARAM_COUNT                                            = 65614
	LIBFPTR_PARAM_COUNTER_TYPE                                     = 65615
	LIBFPTR_PARAM_STEP_COUNTER_TYPE                                = 65616
	LIBFPTR_PARAM_ERROR_TAG_NUMBER                                 = 65617
	LIBFPTR_PARAM_TABLE                                            = 65618
	LIBFPTR_PARAM_ROW                                              = 65619
	LIBFPTR_PARAM_FIELD                                            = 65620
	LIBFPTR_PARAM_FIELD_VALUE                                      = 65621
	LIBFPTR_PARAM_FN_DATA_TYPE                                     = 65622
	LIBFPTR_PARAM_TAG_NUMBER                                       = 65623
	LIBFPTR_PARAM_TAG_VALUE                                        = 65624
	LIBFPTR_PARAM_DOCUMENTS_COUNT                                  = 65625
	LIBFPTR_PARAM_FISCAL_SIGN                                      = 65626
	LIBFPTR_PARAM_DEVICE_FFD_VERSION                               = 65627
	LIBFPTR_PARAM_FN_FFD_VERSION                                   = 65628
	LIBFPTR_PARAM_FFD_VERSION                                      = 65629
	LIBFPTR_PARAM_CHECK_SUM                                        = 65630
	LIBFPTR_PARAM_COMMODITY_NAME                                   = 65631
	LIBFPTR_PARAM_PRICE                                            = 65632
	LIBFPTR_PARAM_QUANTITY                                         = 65633
	LIBFPTR_PARAM_POSITION_SUM                                     = 65634
	LIBFPTR_PARAM_FN_TYPE                                          = 65635
	LIBFPTR_PARAM_FN_VERSION                                       = 65636
	LIBFPTR_PARAM_REGISTRATIONS_REMAIN                             = 65637
	LIBFPTR_PARAM_REGISTRATIONS_COUNT                              = 65638
	LIBFPTR_PARAM_NO_ERROR_IF_NOT_SUPPORTED                        = 65639
	LIBFPTR_PARAM_OFD_EXCHANGE_STATUS                              = 65640
	LIBFPTR_PARAM_FN_ERROR_DATA                                    = 65641
	LIBFPTR_PARAM_FN_ERROR_CODE                                    = 65642
	LIBFPTR_PARAM_ENVD_MODE                                        = 65643
	LIBFPTR_PARAM_DOCUMENT_CLOSED                                  = 65644
	LIBFPTR_PARAM_JSON_DATA                                        = 65645
	LIBFPTR_PARAM_COMMAND_SUBSYSTEM                                = 65646
	LIBFPTR_PARAM_FN_OPERATION_TYPE                                = 65647
	LIBFPTR_PARAM_FN_STATE                                         = 65648
	LIBFPTR_PARAM_ENVD_MODE_ENABLED                                = 65649
	LIBFPTR_PARAM_SETTING_ID                                       = 65650
	LIBFPTR_PARAM_SETTING_VALUE                                    = 65651
	LIBFPTR_PARAM_MAPPING_KEY                                      = 65652
	LIBFPTR_PARAM_MAPPING_VALUE                                    = 65653
	LIBFPTR_PARAM_COMMODITY_PIECE                                  = 65654
	LIBFPTR_PARAM_POWER_SOURCE_TYPE                                = 65655
	LIBFPTR_PARAM_BATTERY_CHARGE                                   = 65656
	LIBFPTR_PARAM_VOLTAGE                                          = 65657
	LIBFPTR_PARAM_USE_BATTERY                                      = 65658
	LIBFPTR_PARAM_BATTERY_CHARGING                                 = 65659
	LIBFPTR_PARAM_CAN_PRINT_WHILE_ON_BATTERY                       = 65660
	LIBFPTR_PARAM_MAC_ADDRESS                                      = 65661
	LIBFPTR_PARAM_FN_FISCAL                                        = 65662
	LIBFPTR_PARAM_NETWORK_ERROR                                    = 65663
	LIBFPTR_PARAM_OFD_ERROR                                        = 65664
	LIBFPTR_PARAM_FN_ERROR                                         = 65665
	LIBFPTR_PARAM_COMMAND_CODE                                     = 65666
	LIBFPTR_PARAM_PRINTER_TEMPERATURE                              = 65667
	LIBFPTR_PARAM_RECORDS_TYPE                                     = 65668
	LIBFPTR_PARAM_OFD_FISCAL_SIGN                                  = 65669
	LIBFPTR_PARAM_HAS_OFD_TICKET                                   = 65670
	LIBFPTR_PARAM_NO_SERIAL_NUMBER                                 = 65671
	LIBFPTR_PARAM_RTC_FAULT                                        = 65672
	LIBFPTR_PARAM_SETTINGS_FAULT                                   = 65673
	LIBFPTR_PARAM_COUNTERS_FAULT                                   = 65674
	LIBFPTR_PARAM_USER_MEMORY_FAULT                                = 65675
	LIBFPTR_PARAM_SERVICE_COUNTERS_FAULT                           = 65676
	LIBFPTR_PARAM_ATTRIBUTES_FAULT                                 = 65677
	LIBFPTR_PARAM_FN_FAULT                                         = 65678
	LIBFPTR_PARAM_INVALID_FN                                       = 65679
	LIBFPTR_PARAM_HARD_FAULT                                       = 65680
	LIBFPTR_PARAM_MEMORY_MANAGER_FAULT                             = 65681
	LIBFPTR_PARAM_SCRIPTS_FAULT                                    = 65682
	LIBFPTR_PARAM_FULL_RESET                                       = 65683
	LIBFPTR_PARAM_WAIT_FOR_REBOOT                                  = 65684
	LIBFPTR_PARAM_SCALE_PERCENT                                    = 65685
	LIBFPTR_PARAM_FN_NEED_REPLACEMENT                              = 65686
	LIBFPTR_PARAM_FN_RESOURCE_EXHAUSTED                            = 65687
	LIBFPTR_PARAM_FN_MEMORY_OVERFLOW                               = 65688
	LIBFPTR_PARAM_FN_OFD_TIMEOUT                                   = 65689
	LIBFPTR_PARAM_FN_CRITICAL_ERROR                                = 65690
	LIBFPTR_PARAM_OFD_MESSAGE_READ                                 = 65691
	LIBFPTR_PARAM_DEVICE_MIN_FFD_VERSION                           = 65692
	LIBFPTR_PARAM_DEVICE_MAX_FFD_VERSION                           = 65693
	LIBFPTR_PARAM_DEVICE_UPTIME                                    = 65694
	LIBFPTR_PARAM_NOMENCLATURE_TYPE                                = 65695
	LIBFPTR_PARAM_GTIN                                             = 65696
	LIBFPTR_PARAM_FN_DOCUMENT_TYPE                                 = 65697
	LIBFPTR_PARAM_NETWORK_ERROR_TEXT                               = 65698
	LIBFPTR_PARAM_FN_ERROR_TEXT                                    = 65699
	LIBFPTR_PARAM_OFD_ERROR_TEXT                                   = 65700
	LIBFPTR_PARAM_USER_SCRIPT_ID                                   = 65701
	LIBFPTR_PARAM_USER_SCRIPT_PARAMETER                            = 65702
	LIBFPTR_PARAM_USER_MEMORY_OPERATION                            = 65703
	LIBFPTR_PARAM_USER_MEMORY_DATA                                 = 65704
	LIBFPTR_PARAM_USER_MEMORY_STRING                               = 65705
	LIBFPTR_PARAM_USER_MEMORY_ADDRESS                              = 65706
	LIBFPTR_PARAM_FN_PRESENT                                       = 65707
	LIBFPTR_PARAM_BLOCKED                                          = 65708
	LIBFPTR_PARAM_DOCUMENT_PRINTED                                 = 65709
	LIBFPTR_PARAM_DISCOUNT_SUM                                     = 65710
	LIBFPTR_PARAM_SURCHARGE_SUM                                    = 65711
	LIBFPTR_PARAM_LK_USER_CODE                                     = 65712
	LIBFPTR_PARAM_LICENSE_COUNT                                    = 65713
	LIBFPTR_PARAM_DEFER                                            = 65714
	LIBFPTR_PARAM_CAP_54FZ                                         = 65715
	LIBFPTR_PARAM_CAP_MANUAL_CLICHE_CONTROL                        = 65716
	LIBFPTR_PARAM_CAP_PAYMENTS_COUNT                               = 65717
	LIBFPTR_PARAM_FIRMWARE_CHUNK_SIZE                              = 65718
	LIBFPTR_PARAM_FIRMWARE_CHUNK_DATA                              = 65719
	LIBFPTR_PARAM_FN_FLAGS                                         = 65720
	LIBFPTR_PARAM_PRINT_FOOTER                                     = 65721
	LIBFPTR_PARAM_PUBLIC_KEY                                       = 65722
	LIBFPTR_PARAM_MAGIC_NUMBER                                     = 65723
	LIBFPTR_PARAM_SIGN                                             = 65724
	LIBFPTR_PARAM_SOFT_NAME                                        = 65725
	LIBFPTR_PARAM_SESSION_CODE                                     = 65726
	LIBFPTR_PARAM_ETHERNET_CONFIG_TIMEOUT                          = 65727
	LIBFPTR_PARAM_ETHERNET_DHCP                                    = 65728
	LIBFPTR_PARAM_ETHERNET_IP                                      = 65729
	LIBFPTR_PARAM_ETHERNET_MASK                                    = 65730
	LIBFPTR_PARAM_ETHERNET_GATEWAY                                 = 65731
	LIBFPTR_PARAM_ETHERNET_PORT                                    = 65732
	LIBFPTR_PARAM_ETHERNET_DNS_IP                                  = 65733
	LIBFPTR_PARAM_ETHERNET_DNS_STATIC                              = 65734
	LIBFPTR_PARAM_STORE_IN_JOURNAL                                 = 65735
	LIBFPTR_PARAM_NEW_PLATFORM                                     = 65736
	LIBFPTR_PARAM_UNIT_RELEASE_VERSION                             = 65737
	LIBFPTR_PARAM_USE_VAT18                                        = 65738
	LIBFPTR_PARAM_TAG_NAME                                         = 65739
	LIBFPTR_PARAM_TAG_TYPE                                         = 65740
	LIBFPTR_PARAM_TAG_IS_COMPLEX                                   = 65741
	LIBFPTR_PARAM_TAG_IS_REPEATABLE                                = 65742
	LIBFPTR_PARAM_SHIFT_AUTO_OPENED                                = 65743
	LIBFPTR_PARAM_CONTAINER_FIRMWARE_VERSION                       = 65744
	LIBFPTR_PARAM_CONTAINER_CONFIGURATION_VERSION                  = 65745
	LIBFPTR_PARAM_CONTAINER_BOOTLOADER_VERSION                     = 65746
	LIBFPTR_PARAM_CONTAINER_SCRIPTS_VERSION                        = 65747
	LIBFPTR_PARAM_PAPER_NEAR_END                                   = 65748
	LIBFPTR_PARAM_REPORT_ELECTRONICALLY                            = 65749
	LIBFPTR_PARAM_ACTIVATION_METHOD                                = 65750
	LIBFPTR_PARAM_KEYS                                             = 65751
	LIBFPTR_PARAM_UIN                                              = 65752
	LIBFPTR_PARAM_VERSION                                          = 65753
	LIBFPTR_PARAM_PUBLIC_KEY_SIGN                                  = 65754
	LIBFPTR_PARAM_CAP_DISABLE_PRINT_REPORTS                        = 65755
	LIBFPTR_PARAM_REGISTRATION_NUMBER                              = 65756
	LIBFPTR_PARAM_PIXEL_BUFFER                                     = 65757
	LIBFPTR_PARAM_REPEAT_NUMBER                                    = 65758
	LIBFPTR_PARAM_FIELD_TYPE                                       = 65759
	LIBFPTR_PARAM_MARKING_CODE                                     = 65760
	LIBFPTR_PARAM_CONTAINER_DIRECT_BOOT_VERSION                    = 65761
	LIBFPTR_PARAM_SCRIPT_NAME                                      = 65762
	LIBFPTR_PARAM_SCRIPT_HASH                                      = 65763
	LIBFPTR_PARAM_RECORDS_ID                                       = 65764
	LIBFPTR_PARAM_USER_SCRIPT_RESULT_1                             = 65765
	LIBFPTR_PARAM_USER_SCRIPT_RESULT_2                             = 65766
	LIBFPTR_PARAM_USER_SCRIPT_RESULT_3                             = 65767
	LIBFPTR_PARAM_USER_SCRIPT_RESULT_4                             = 65768
	LIBFPTR_PARAM_USER_SCRIPT_RESULT_5                             = 65769
	LIBFPTR_PARAM_IS_USER_SCRIPT                                   = 65770
	LIBFPTR_PARAM_DOCUMENT_NUMBER_END                              = 65771
	LIBFPTR_PARAM_SHIFT_NUMBER_END                                 = 65772
	LIBFPTR_PARAM_SCRIPT_CODE                                      = 65773
	LIBFPTR_PARAM_SCRIPT_RESULT                                    = 65774
	LIBFPTR_PARAM_SCRIPT_TYPE                                      = 65775
	LIBFPTR_PARAM_WIFI_CONFIG_TIMEOUT                              = 65776
	LIBFPTR_PARAM_WIFI_DHCP                                        = 65777
	LIBFPTR_PARAM_WIFI_IP                                          = 65778
	LIBFPTR_PARAM_WIFI_MASK                                        = 65779
	LIBFPTR_PARAM_WIFI_GATEWAY                                     = 65780
	LIBFPTR_PARAM_WIFI_PORT                                        = 65781
	LIBFPTR_PARAM_UC_VERSION                                       = 65782
	LIBFPTR_PARAM_UC_AVAILABLE_MEMORY                              = 65783
	LIBFPTR_PARAM_UC_USED_MEMORY_BY_SUMS                           = 65784
	LIBFPTR_PARAM_UC_USED_MEMORY_BY_QUANTITIES                     = 65785
	LIBFPTR_PARAM_UC_LAYER_1                                       = 65786
	LIBFPTR_PARAM_UC_FLAGS_1                                       = 65787
	LIBFPTR_PARAM_UC_MASK_1                                        = 65788
	LIBFPTR_PARAM_UC_LAYER_VALUE_1                                 = 65789
	LIBFPTR_PARAM_UC_LAYER_2                                       = 65790
	LIBFPTR_PARAM_UC_FLAGS_2                                       = 65791
	LIBFPTR_PARAM_UC_MASK_2                                        = 65792
	LIBFPTR_PARAM_UC_LAYER_VALUE_2                                 = 65793
	LIBFPTR_PARAM_UC_LAYER_3                                       = 65794
	LIBFPTR_PARAM_UC_FLAGS_3                                       = 65795
	LIBFPTR_PARAM_UC_MASK_3                                        = 65796
	LIBFPTR_PARAM_UC_LAYER_VALUE_3                                 = 65797
	LIBFPTR_PARAM_UC_LAYER_4                                       = 65798
	LIBFPTR_PARAM_UC_FLAGS_4                                       = 65799
	LIBFPTR_PARAM_UC_MASK_4                                        = 65800
	LIBFPTR_PARAM_UC_LAYER_VALUE_4                                 = 65801
	LIBFPTR_PARAM_RECEIPTS_COUNT                                   = 65802
	LIBFPTR_PARAM_PAYMENTS_SUM_CASH                                = 65803
	LIBFPTR_PARAM_PAYMENTS_SUM_ELECTRONICALLY                      = 65804
	LIBFPTR_PARAM_PAYMENTS_SUM_PREPAID                             = 65805
	LIBFPTR_PARAM_PAYMENTS_SUM_CREDIT                              = 65806
	LIBFPTR_PARAM_PAYMENTS_SUM_OTHER                               = 65807
	LIBFPTR_PARAM_TAXES_SUM_VAT20                                  = 65808
	LIBFPTR_PARAM_TAXES_SUM_VAT120                                 = 65809
	LIBFPTR_PARAM_TAXES_SUM_VAT10                                  = 65810
	LIBFPTR_PARAM_TAXES_SUM_VAT110                                 = 65811
	LIBFPTR_PARAM_TAXES_SUM_VAT0                                   = 65812
	LIBFPTR_PARAM_TAXES_SUM_NO                                     = 65813
	LIBFPTR_PARAM_CORRECTIONS_COUNT                                = 65814
	LIBFPTR_PARAM_CORRECTIONS_SUM                                  = 65815
	LIBFPTR_PARAM_FN_COUNTERS_TYPE                                 = 65816
	LIBFPTR_PARAM_FN_DAYS_REMAIN                                   = 65817
	LIBFPTR_PARAM_FREE_MEMORY                                      = 65818
	LIBFPTR_PARAM_FN_MAX_FFD_VERSION                               = 65819
	LIBFPTR_PARAM_RECEIPTS_SUM                                     = 65820
	LIBFPTR_PARAM_LICENSE_NAME                                     = 65821
	LIBFPTR_PARAM_UNIVERSAL_COUNTERS_FAULT                         = 65822
	LIBFPTR_PARAM_USE_LICENSES                                     = 65823
	LIBFPTR_PARAM_LICENSE_VALID_FROM                               = 65824
	LIBFPTR_PARAM_LICENSE_VALID_UNTIL                              = 65825
	LIBFPTR_PARAM_MARKING_CODE_TYPE                                = 65826
	LIBFPTR_PARAM_SETTING_NAME                                     = 65827
	LIBFPTR_PARAM_SETTING_TYPE                                     = 65828
	LIBFPTR_PARAM_FONT_WIDTH                                       = 65829
	LIBFPTR_PARAM_REMOTE_CALL                                      = 65830
	LIBFPTR_PARAM_SCRIPT_PARAMS                                    = 65831
	LIBFPTR_PARAM_IGNORE_EMPTY                                     = 65832
	LIBFPTR_PARAM_METHOD_DATA                                      = 65833
	LIBFPTR_PARAM_METHOD_RESULT                                    = 65834
	LIBFPTR_PARAM_RPC_SERVER_OS                                    = 65835
	LIBFPTR_PARAM_RPC_SERVER_VERSION                               = 65836
	LIBFPTR_PARAM_RPC_DRIVER_VERSION                               = 65837
	LIBFPTR_PARAM_LOCKED                                           = 65838
	LIBFPTR_PARAM_BOUND                                            = 65839
	LIBFPTR_PARAM_COMMODITIES_TABLE_FAULT                          = 65840
	LIBFPTR_PARAM_HAS_ADDITIONAL_DATA                              = 65841
	LIBFPTR_PARAM_FISCAL_SIGN_ARCHIVE                              = 65842
	LIBFPTR_PARAM_COMMAND_GROUP                                    = 65843
	LIBFPTR_PARAM_ERROR_CODE                                       = 65844
	LIBFPTR_PARAM_MARKING_WAIT_FOR_VALIDATION_RESULT               = 65845
	LIBFPTR_PARAM_MARKING_CODE_STATUS                              = 65846
	LIBFPTR_PARAM_MARKING_CODE_VALIDATION_RESULT                   = 65847
	LIBFPTR_PARAM_MARKING_CODE_OFFLINE_VALIDATION_ERROR            = 65848
	LIBFPTR_PARAM_MARKING_CODE_ONLINE_VALIDATION_ERROR             = 65849
	LIBFPTR_PARAM_MARKING_CODE_VALIDATION_READY                    = 65850
	LIBFPTR_PARAM_MEASUREMENT_UNIT                                 = 65851
	LIBFPTR_PARAM_MARKING_PROCESSING_MODE                          = 65852
	LIBFPTR_PARAM_MARKING_FRACTIONAL_QUANTITY                      = 65853
	LIBFPTR_PARAM_PRODUCT_CODE                                     = 65854
	LIBFPTR_PARAM_TRADE_MARKED_PRODUCTS                            = 65855
	LIBFPTR_PARAM_INSURANCE_ACTIVITY                               = 65856
	LIBFPTR_PARAM_PAWN_SHOP_ACTIVITY                               = 65857
	LIBFPTR_PARAM_TLV_LIST                                         = 65858
	LIBFPTR_PARAM_CHECK_MARKING_SERVER_READY                       = 65859
	LIBFPTR_PARAM_MARKING_SERVER_RESPONSE_TIME                     = 65860
	LIBFPTR_PARAM_MARKING_SERVER_ERROR_CODE                        = 65861
	LIBFPTR_PARAM_MARKING_SERVER_ERROR_DESCRIPTION                 = 65862
	LIBFPTR_PARAM_ISM_ERROR                                        = 65863
	LIBFPTR_PARAM_ISM_ERROR_TEXT                                   = 65864
	LIBFPTR_PARAM_MARKING_MODE_CHECKING_STATUS                     = 65865
	LIBFPTR_PARAM_MARK_CHECKING_COUNT                              = 65866
	LIBFPTR_PARAM_MARK_SOLD_COUNT                                  = 65867
	LIBFPTR_PARAM_NOTICE_IS_BEGIN                                  = 65868
	LIBFPTR_PARAM_NOTICE_FREE_MEMORY                               = 65869
	LIBFPTR_PARAM_NOTICE_COUNT                                     = 65870
	LIBFPTR_PARAM_MARKING_NOT_SEND_TO_SERVER                       = 65871
	LIBFPTR_PARAM_DOCUMENT_TYPE                                    = 65872
	LIBFPTR_PARAM_PRINT_REPORT                                     = 65873
	LIBFPTR_PARAM_FN_EXECUTION                                     = 65874
	LIBFPTR_PARAM_MCU_SN                                           = 65875
	LIBFPTR_PARAM_MCU_PART_ID                                      = 65876
	LIBFPTR_PARAM_MCU_PART_NAME                                    = 65877
	LIBFPTR_PARAM_IS_REQUEST_SENT                                  = 65878
	LIBFPTR_PARAM_FN_CHECK_MARK_TIME                               = 65879
	LIBFPTR_PARAM_SENDING_MARK_TIME                                = 65880
	LIBFPTR_PARAM_MARKING_SERVER_EXCHANGE_TIME                     = 65881
	LIBFPTR_PARAM_FULL_SENDING_MARK_TIME                           = 65882
	LIBFPTR_PARAM_MARK_CHECKING_STATUS_IN_CASH                     = 65883
	LIBFPTR_PARAM_MARK_CHECKING_TYPE_IN_CASH                       = 65884
	LIBFPTR_PARAM_MARK_CHECKING_STAGE_IN_CASH                      = 65885
	LIBFPTR_PARAM_MARKING_CODE_ONLINE_VALIDATION_RESULT            = 65886
	LIBFPTR_PARAM_MARKING_CODE_ONLINE_VALIDATION_ERROR_DESCRIPTION = 65887
	LIBFPTR_PARAM_FN_CONTAINS_KEYS_UPDATER_SERVER_URI              = 65888
	LIBFPTR_PARAM_MARKING_CODE_CLEAR                               = 65889
	LIBFPTR_PARAM_MODULE_ADDRESS                                   = 65890
	LIBFPTR_PARAM_SEGMENT_ADDRESS                                  = 65891
	LIBFPTR_PARAM_LAST_SUCCESSFUL_OKP                              = 65892
	LIBFPTR_PARAM_FN_SERIAL_NUMBER                                 = 65893
	LIBFPTR_PARAM_ECR_REGISTRATION_NUMBER                          = 65894
	LIBFPTR_PARAM_OFD_VATIN                                        = 65895
	LIBFPTR_PARAM_FNS_URL                                          = 65896
	LIBFPTR_PARAM_MACHINE_NUMBER                                   = 65897
	LIBFPTR_PARAM_MARKING_PRODUCT_ID                               = 65898
	LIBFPTR_PARAM_TIMEOUT                                          = 65899
	LIBFPTR_PARAM_PRINT_UPDATE_FNM_KEYS_REPORT                     = 65900
	LIBFPTR_PARAM_FN_KEYS_UPDATER_SERVER_URI                       = 65901
	LIBFPTR_PARAM_DOCUMENT_ELECTRONICALLY                          = 65902
	LIBFPTR_PARAM_FORMAT_TEXT                                      = 65903
	LIBFPTR_PARAM_RECEIPT_SIZE                                     = 65904
	LIBFPTR_PARAM_MARK_SIZE                                        = 65905
	LIBFPTR_PARAM_MCU_TEMPERATURE                                  = 65906
	LIBFPTR_PARAM_DATA_FOR_SEND_IS_EMPTY                           = 65907
	LIBFPTR_PARAM_AVAILABLE_CLOSING                                = 65908
	LIBFPTR_PARAM_AVAILABLE_CANCELLATION                           = 65909
	LIBFPTR_PARAM_AVAILABLE_POSITION_ADDING                        = 65910
	LIBFPTR_PARAM_AVAILABLE_PAYMENT                                = 65911
	LIBFPTR_PARAM_AVAILABLE_TOTAL                                  = 65912
	LIBFPTR_PARAM_AVAILABLE_ATTRIBUTES_ADDING                      = 65913
	LIBFPTR_PARAM_OPERATOR_REGISTERED                              = 65914
	LIBFPTR_PARAM_DEVICE_PLATFORM_VERSION                          = 65915
	LIBFPTR_PARAM_GUID                                             = 65916
	LIBFPTR_PARAM_PATTERN_REGISTERS                                = 65917
	LIBFPTR_PARAM_PATTERN_TAGS                                     = 65918
	LIBFPTR_PARAM_PATTERN_SETTINGS                                 = 65919
)

const (
	LIBFPTR_OK                                            = 0
	LIBFPTR_ERROR_CONNECTION_DISABLED                     = 1
	LIBFPTR_ERROR_NO_CONNECTION                           = 2
	LIBFPTR_ERROR_PORT_BUSY                               = 3
	LIBFPTR_ERROR_PORT_NOT_AVAILABLE                      = 4
	LIBFPTR_ERROR_INCORRECT_DATA                          = 5
	LIBFPTR_ERROR_INTERNAL                                = 6
	LIBFPTR_ERROR_UNSUPPORTED_CAST                        = 7
	LIBFPTR_ERROR_NO_REQUIRED_PARAM                       = 8
	LIBFPTR_ERROR_INVALID_SETTINGS                        = 9
	LIBFPTR_ERROR_NOT_CONFIGURED                          = 10
	LIBFPTR_ERROR_NOT_SUPPORTED                           = 11
	LIBFPTR_ERROR_INVALID_MODE                            = 12
	LIBFPTR_ERROR_INVALID_PARAM                           = 13
	LIBFPTR_ERROR_NOT_LOADED                              = 14
	LIBFPTR_ERROR_UNKNOWN                                 = 15
	LIBFPTR_ERROR_INVALID_SUM                             = 16
	LIBFPTR_ERROR_INVALID_QUANTITY                        = 17
	LIBFPTR_ERROR_CASH_COUNTER_OVERFLOW                   = 18
	LIBFPTR_ERROR_LAST_OPERATION_STORNO_DENIED            = 19
	LIBFPTR_ERROR_STORNO_BY_CODE_DENIED                   = 20
	LIBFPTR_ERROR_LAST_OPERATION_NOT_REPEATABLE           = 21
	LIBFPTR_ERROR_DISCOUNT_NOT_REPEATABLE                 = 22
	LIBFPTR_ERROR_DISCOUNT_DENIED                         = 23
	LIBFPTR_ERROR_INVALID_COMMODITY_CODE                  = 24
	LIBFPTR_ERROR_INVALID_COMMODITY_BARCODE               = 25
	LIBFPTR_ERROR_INVALID_COMMAND_FORMAT                  = 26
	LIBFPTR_ERROR_INVALID_COMMAND_LENGTH                  = 27
	LIBFPTR_ERROR_BLOCKED_IN_DATE_INPUT_MODE              = 28
	LIBFPTR_ERROR_NEED_DATE_ACCEPT                        = 29
	LIBFPTR_ERROR_NO_MORE_DATA                            = 30
	LIBFPTR_ERROR_NO_ACCEPT_OR_CANCEL                     = 31
	LIBFPTR_ERROR_BLOCKED_BY_REPORT_INTERRUPTION          = 32
	LIBFPTR_ERROR_DISABLE_CASH_CONTROL_DENIED             = 33
	LIBFPTR_ERROR_MODE_BLOCKED                            = 34
	LIBFPTR_ERROR_CHECK_DATE_TIME                         = 35
	LIBFPTR_ERROR_DATE_TIME_LESS_THAN_FS                  = 36
	LIBFPTR_ERROR_CLOSE_ARCHIVE_DENIED                    = 37
	LIBFPTR_ERROR_COMMODITY_NOT_FOUND                     = 38
	LIBFPTR_ERROR_WEIGHT_BARCODE_WITH_INVALID_QUANTITY    = 39
	LIBFPTR_ERROR_RECEIPT_BUFFER_OVERFLOW                 = 40
	LIBFPTR_ERROR_QUANTITY_TOO_FEW                        = 41
	LIBFPTR_ERROR_STORNO_TOO_MUCH                         = 42
	LIBFPTR_ERROR_BLOCKED_COMMODITY_NOT_FOUND             = 43
	LIBFPTR_ERROR_NO_PAPER                                = 44
	LIBFPTR_ERROR_COVER_OPENED                            = 45
	LIBFPTR_ERROR_PRINTER_FAULT                           = 46
	LIBFPTR_ERROR_MECHANICAL_FAULT                        = 47
	LIBFPTR_ERROR_INVALID_RECEIPT_TYPE                    = 48
	LIBFPTR_ERROR_INVALID_UNIT_TYPE                       = 49
	LIBFPTR_ERROR_NO_MEMORY                               = 50
	LIBFPTR_ERROR_PICTURE_NOT_FOUND                       = 51
	LIBFPTR_ERROR_NONCACH_PAYMENTS_TOO_MUCH               = 52
	LIBFPTR_ERROR_RETURN_DENIED                           = 53
	LIBFPTR_ERROR_PAYMENTS_OVERFLOW                       = 54
	LIBFPTR_ERROR_BUSY                                    = 55
	LIBFPTR_ERROR_GSM                                     = 56
	LIBFPTR_ERROR_INVALID_DISCOUNT                        = 57
	LIBFPTR_ERROR_OPERATION_AFTER_DISCOUNT_DENIED         = 58
	LIBFPTR_ERROR_INVALID_DEPARTMENT                      = 59
	LIBFPTR_ERROR_INVALID_PAYMENT_TYPE                    = 60
	LIBFPTR_ERROR_MULTIPLICATION_OVERFLOW                 = 61
	LIBFPTR_ERROR_DENIED_BY_SETTINGS                      = 62
	LIBFPTR_ERROR_TOTAL_OVERFLOW                          = 63
	LIBFPTR_ERROR_DENIED_IN_ANNULATION_RECEIPT            = 64
	LIBFPTR_ERROR_JOURNAL_OVERFLOW                        = 65
	LIBFPTR_ERROR_NOT_FULLY_PAID                          = 66
	LIBFPTR_ERROR_DENIED_IN_RETURN_RECEIPT                = 67
	LIBFPTR_ERROR_SHIFT_EXPIRED                           = 68
	LIBFPTR_ERROR_DENIED_IN_SELL_RECEIPT                  = 69
	LIBFPTR_ERROR_FISCAL_MEMORY_OVERFLOW                  = 70
	LIBFPTR_ERROR_INVALID_PASSWORD                        = 71
	LIBFPTR_ERROR_JOURNAL_BUSY                            = 72
	LIBFPTR_ERROR_DENIED_IN_CLOSED_SHIFT                  = 73
	LIBFPTR_ERROR_INVALID_TABLE_NUMBER                    = 74
	LIBFPTR_ERROR_INVALID_ROW_NUMBER                      = 75
	LIBFPTR_ERROR_INVALID_FIELD_NUMBER                    = 76
	LIBFPTR_ERROR_INVALID_DATE_TIME                       = 77
	LIBFPTR_ERROR_INVALID_STORNO_SUM                      = 78
	LIBFPTR_ERROR_CHANGE_CALCULATION                      = 79
	LIBFPTR_ERROR_NO_CASH                                 = 80
	LIBFPTR_ERROR_DENIED_IN_CLOSED_RECEIPT                = 81
	LIBFPTR_ERROR_DENIED_IN_OPENED_RECEIPT                = 82
	LIBFPTR_ERROR_DENIED_IN_OPENED_SHIFT                  = 83
	LIBFPTR_ERROR_SERIAL_NUMBER_ALREADY_ENTERED           = 84
	LIBFPTR_ERROR_TOO_MUCH_REREGISTRATIONS                = 85
	LIBFPTR_ERROR_INVALID_SHIFT_NUMBER                    = 86
	LIBFPTR_ERROR_INVALID_SERIAL_NUMBER                   = 87
	LIBFPTR_ERROR_INVALID_RNM_VATIN                       = 88
	LIBFPTR_ERROR_FISCAL_PRINTER_NOT_ACTIVATED            = 89
	LIBFPTR_ERROR_SERIAL_NUMBER_NOT_ENTERED               = 90
	LIBFPTR_ERROR_NO_MORE_REPORTS                         = 91
	LIBFPTR_ERROR_MODE_NOT_ACTIVATED                      = 92
	LIBFPTR_ERROR_RECORD_NOT_FOUND_IN_JOURNAL             = 93
	LIBFPTR_ERROR_INVALID_LICENSE                         = 94
	LIBFPTR_ERROR_NEED_FULL_RESET                         = 95
	LIBFPTR_ERROR_DENIED_BY_LICENSE                       = 96
	LIBFPTR_ERROR_DISCOUNT_CANCELLATION_DENIED            = 97
	LIBFPTR_ERROR_CLOSE_RECEIPT_DENIED                    = 98
	LIBFPTR_ERROR_INVALID_ROUTE_NUMBER                    = 99
	LIBFPTR_ERROR_INVALID_START_ZONE_NUMBER               = 100
	LIBFPTR_ERROR_INVALID_END_ZONE_NUMBER                 = 101
	LIBFPTR_ERROR_INVALID_RATE_TYPE                       = 102
	LIBFPTR_ERROR_INVALID_RATE                            = 103
	LIBFPTR_ERROR_FISCAL_MODULE_EXCHANGE                  = 104
	LIBFPTR_ERROR_NEED_TECHNICAL_SUPPORT                  = 105
	LIBFPTR_ERROR_SHIFT_NUMBERS_DID_NOT_MATCH             = 106
	LIBFPTR_ERROR_DEVICE_NOT_FOUND                        = 107
	LIBFPTR_ERROR_EXTERNAL_DEVICE_CONNECTION              = 108
	LIBFPTR_ERROR_DISPENSER_INVALID_STATE                 = 109
	LIBFPTR_ERROR_INVALID_POSITIONS_COUNT                 = 110
	LIBFPTR_ERROR_DISPENSER_INVALID_NUMBER                = 111
	LIBFPTR_ERROR_INVALID_DIVIDER                         = 112
	LIBFPTR_ERROR_FN_ACTIVATION_DENIED                    = 113
	LIBFPTR_ERROR_PRINTER_OVERHEAT                        = 114
	LIBFPTR_ERROR_FN_EXCHANGE                             = 115
	LIBFPTR_ERROR_FN_INVALID_FORMAT                       = 116
	LIBFPTR_ERROR_FN_INVALID_STATE                        = 117
	LIBFPTR_ERROR_FN_FAULT                                = 118
	LIBFPTR_ERROR_FN_CRYPTO_FAULT                         = 119
	LIBFPTR_ERROR_FN_EXPIRED                              = 120
	LIBFPTR_ERROR_FN_OVERFLOW                             = 121
	LIBFPTR_ERROR_FN_INVALID_DATE_TIME                    = 122
	LIBFPTR_ERROR_FN_NO_MORE_DATA                         = 123
	LIBFPTR_ERROR_FN_TOTAL_OVERFLOW                       = 124
	LIBFPTR_ERROR_BUFFER_OVERFLOW                         = 125
	LIBFPTR_ERROR_PRINT_SECOND_COPY_DENIED                = 126
	LIBFPTR_ERROR_NEED_RESET_JOURNAL                      = 127
	LIBFPTR_ERROR_TAX_SUM_TOO_MUCH                        = 128
	LIBFPTR_ERROR_TAX_ON_LAST_OPERATION_DENIED            = 129
	LIBFPTR_ERROR_INVALID_FN_NUMBER                       = 130
	LIBFPTR_ERROR_TAX_CANCEL_DENIED                       = 131
	LIBFPTR_ERROR_LOW_BATTERY                             = 132
	LIBFPTR_ERROR_FN_INVALID_COMMAND                      = 133
	LIBFPTR_ERROR_FN_COMMAND_OVERFLOW                     = 134
	LIBFPTR_ERROR_FN_NO_TRANSPORT_CONNECTION              = 135
	LIBFPTR_ERROR_FN_CRYPTO_HAS_EXPIRED                   = 136
	LIBFPTR_ERROR_FN_RESOURCE_HAS_EXPIRED                 = 137
	LIBFPTR_ERROR_INVALID_MESSAGE_FROM_OFD                = 138
	LIBFPTR_ERROR_FN_HAS_NOT_SEND_DOCUMENTS               = 139
	LIBFPTR_ERROR_FN_TIMEOUT                              = 140
	LIBFPTR_ERROR_FN_SHIFT_EXPIRED                        = 141
	LIBFPTR_ERROR_FN_INVALID_TIME_DIFFERENCE              = 142
	LIBFPTR_ERROR_INVALID_TAXATION_TYPE                   = 143
	LIBFPTR_ERROR_INVALID_TAX_TYPE                        = 144
	LIBFPTR_ERROR_INVALID_COMMODITY_PAYMENT_TYPE          = 145
	LIBFPTR_ERROR_INVALID_COMMODITY_CODE_TYPE             = 146
	LIBFPTR_ERROR_EXCISABLE_COMMODITY_DENIED              = 147
	LIBFPTR_ERROR_FISCAL_PROPERTY_WRITE                   = 148
	LIBFPTR_ERROR_INVALID_COUNTER_TYPE                    = 149
	LIBFPTR_ERROR_CUTTER_FAULT                            = 150
	LIBFPTR_ERROR_REPORT_INTERRUPTED                      = 151
	LIBFPTR_ERROR_INVALID_LEFT_MARGIN                     = 152
	LIBFPTR_ERROR_INVALID_ALIGNMENT                       = 153
	LIBFPTR_ERROR_INVALID_TAX_MODE                        = 154
	LIBFPTR_ERROR_FILE_NOT_FOUND                          = 155
	LIBFPTR_ERROR_PICTURE_TOO_BIG                         = 156
	LIBFPTR_ERROR_INVALID_BARCODE_PARAMS                  = 157
	LIBFPTR_ERROR_FISCAL_PROPERTY_DENIED                  = 158
	LIBFPTR_ERROR_FN_INTERFACE                            = 159
	LIBFPTR_ERROR_DATA_DUPLICATE                          = 160
	LIBFPTR_ERROR_NO_REQUIRED_FISCAL_PROPERTY             = 161
	LIBFPTR_ERROR_FN_READ_DOCUMENT                        = 162
	LIBFPTR_ERROR_FLOAT_OVERFLOW                          = 163
	LIBFPTR_ERROR_INVALID_SETTING_VALUE                   = 164
	LIBFPTR_ERROR_HARD_FAULT                              = 165
	LIBFPTR_ERROR_FN_NOT_FOUND                            = 166
	LIBFPTR_ERROR_INVALID_AGENT_FISCAL_PROPERTY           = 167
	LIBFPTR_ERROR_INVALID_FISCAL_PROPERTY_VALUE_1002_1056 = 168
	LIBFPTR_ERROR_INVALID_FISCAL_PROPERTY_VALUE_1002_1017 = 169
	LIBFPTR_ERROR_SCRIPT                                  = 170
	LIBFPTR_ERROR_INVALID_USER_MEMORY_INDEX               = 171
	LIBFPTR_ERROR_NO_ACTIVE_OPERATOR                      = 172
	LIBFPTR_ERROR_REGISTRATION_REPORT_INTERRUPTED         = 173
	LIBFPTR_ERROR_CLOSE_FN_REPORT_INTERRUPTED             = 174
	LIBFPTR_ERROR_OPEN_SHIFT_REPORT_INTERRUPTED           = 175
	LIBFPTR_ERROR_OFD_EXCHANGE_REPORT_INTERRUPTED         = 176
	LIBFPTR_ERROR_CLOSE_RECEIPT_INTERRUPTED               = 177
	LIBFPTR_ERROR_FN_QUERY_INTERRUPTED                    = 178
	LIBFPTR_ERROR_RTC_FAULT                               = 179
	LIBFPTR_ERROR_MEMORY_FAULT                            = 180
	LIBFPTR_ERROR_CHIP_FAULT                              = 181
	LIBFPTR_ERROR_TEMPLATES_CORRUPTED                     = 182
	LIBFPTR_ERROR_INVALID_MAC_ADDRESS                     = 183
	LIBFPTR_ERROR_INVALID_SCRIPT_NUMBER                   = 184
	LIBFPTR_ERROR_SCRIPTS_FAULT                           = 185
	LIBFPTR_ERROR_INVALID_SCRIPTS_VERSION                 = 186
	LIBFPTR_ERROR_INVALID_CLICHE_FORMAT                   = 187
	LIBFPTR_ERROR_WAIT_FOR_REBOOT                         = 188
	LIBFPTR_ERROR_NO_LICENSE                              = 189
	LIBFPTR_ERROR_INVALID_FFD_VERSION                     = 190
	LIBFPTR_ERROR_CHANGE_SETTING_DENIED                   = 191
	LIBFPTR_ERROR_INVALID_NOMENCLATURE_TYPE               = 192
	LIBFPTR_ERROR_INVALID_GTIN                            = 193
	LIBFPTR_ERROR_NEGATIVE_MATH_RESULT                    = 194
	LIBFPTR_ERROR_FISCAL_PROPERTIES_COMBINATION           = 195
	LIBFPTR_ERROR_OPERATOR_LOGIN                          = 196
	LIBFPTR_ERROR_INVALID_INTERNET_CHANNEL                = 197
	LIBFPTR_ERROR_DATETIME_NOT_SYNCRONIZED                = 198
	LIBFPTR_ERROR_JOURNAL                                 = 199
	LIBFPTR_ERROR_DENIED_IN_OPENED_DOC                    = 200
	LIBFPTR_ERROR_DENIED_IN_CLOSED_DOC                    = 201
	LIBFPTR_ERROR_LICENSE_MEMORY_OVERFLOW                 = 202
	LIBFPTR_ERROR_NEED_CANCEL_DOCUMENT                    = 203
	LIBFPTR_ERROR_REGISTERS_NOT_INITIALIZED               = 204
	LIBFPTR_ERROR_TOTAL_REQUIRED                          = 205
	LIBFPTR_ERROR_SETTINGS_FAULT                          = 206
	LIBFPTR_ERROR_COUNTERS_FAULT                          = 207
	LIBFPTR_ERROR_USER_MEMORY_FAULT                       = 208
	LIBFPTR_ERROR_SERVICE_COUNTERS_FAULT                  = 209
	LIBFPTR_ERROR_ATTRIBUTES_FAULT                        = 210
	LIBFPTR_ERROR_ALREADY_IN_UPDATE_MODE                  = 211
	LIBFPTR_ERROR_INVALID_FIRMWARE                        = 212
	LIBFPTR_ERROR_INVALID_CHANNEL                         = 213
	LIBFPTR_ERROR_INTERFACE_DOWN                          = 214
	LIBFPTR_ERROR_INVALID_FISCAL_PROPERTY_VALUE_1212_1030 = 215
	LIBFPTR_ERROR_INVALID_FISCAL_PROPERTY_VALUE_1214      = 216
	LIBFPTR_ERROR_INVALID_FISCAL_PROPERTY_VALUE_1212      = 217
	LIBFPTR_ERROR_SYNC_TIME                               = 218
	LIBFPTR_ERROR_VAT18_VAT20_IN_RECEIPT                  = 219
	LIBFPTR_ERROR_PICTURE_NOT_CLOSED                      = 220
	LIBFPTR_ERROR_INTERFACE_BUSY                          = 221
	LIBFPTR_ERROR_INVALID_PICTURE_NUMBER                  = 222
	LIBFPTR_ERROR_INVALID_CONTAINER                       = 223
	LIBFPTR_ERROR_ARCHIVE_CLOSED                          = 224
	LIBFPTR_ERROR_NEED_REGISTRATION                       = 225
	LIBFPTR_ERROR_DENIED_DURING_UPDATE                    = 226
	LIBFPTR_ERROR_INVALID_TOTAL                           = 227
	LIBFPTR_ERROR_MARKING_CODE_CONFLICT                   = 228
	LIBFPTR_ERROR_INVALID_RECORDS_ID                      = 229
	LIBFPTR_ERROR_INVALID_SIGNATURE                       = 230
	LIBFPTR_ERROR_INVALID_EXCISE_SUM                      = 231
	LIBFPTR_ERROR_NO_DOCUMENTS_FOUND_IN_JOURNAL           = 232
	LIBFPTR_ERROR_INVALID_SCRIPT_TYPE                     = 233
	LIBFPTR_ERROR_INVALID_SCRIPT_NAME                     = 234
	LIBFPTR_ERROR_INVALID_POSITIONS_COUNT_WITH_1162       = 235
	LIBFPTR_ERROR_INVALID_UC_COUNTER                      = 236
	LIBFPTR_ERROR_INVALID_UC_TAG                          = 237
	LIBFPTR_ERROR_INVALID_UC_IDX                          = 238
	LIBFPTR_ERROR_INVALID_UC_SIZE                         = 239
	LIBFPTR_ERROR_INVALID_UC_CONFIG                       = 240
	LIBFPTR_ERROR_CONNECTION_LOST                         = 241
	LIBFPTR_ERROR_UNIVERSAL_COUNTERS_FAULT                = 242
	LIBFPTR_ERROR_INVALID_TAX_SUM                         = 243
	LIBFPTR_ERROR_INVALID_MARKING_CODE_TYPE               = 244
	LIBFPTR_ERROR_LICENSE_HARD_FAULT                      = 245
	LIBFPTR_ERROR_LICENSE_INVALID_SIGN                    = 246
	LIBFPTR_ERROR_LICENSE_INVALID_SERIAL                  = 247
	LIBFPTR_ERROR_LICENSE_INVALID_TIME                    = 248
	LIBFPTR_ERROR_DOCUMENT_CANCELED                       = 249
	LIBFPTR_ERROR_INVALID_SCRIPT_PARAMS                   = 250
	LIBFPTR_ERROR_CLICHE_TOO_LONG                         = 251
	LIBFPTR_ERROR_COMMODITIES_TABLE_FAULT                 = 252
	LIBFPTR_ERROR_COMMODITIES_TABLE                       = 253
	LIBFPTR_ERROR_COMMODITIES_TABLE_INVALID_TAG           = 254
	LIBFPTR_ERROR_COMMODITIES_TABLE_INVALID_TAG_SIZE      = 255
	LIBFPTR_ERROR_COMMODITIES_TABLE_NO_TAG_DATA           = 256
	LIBFPTR_ERROR_COMMODITIES_TABLE_NO_FREE_MEMORY        = 257
	LIBFPTR_ERROR_INVALID_CACHE                           = 258
	LIBFPTR_ERROR_SCHEDULER_NOT_READY                     = 259
	LIBFPTR_ERROR_SCHEDULER_INVALID_TASK                  = 260
	LIBFPTR_ERROR_MINIPOS_NO_POSITION_PAYMENT             = 261
	LIBFPTR_ERROR_MINIPOS_COMMAND_TIME_OUT                = 262
	LIBFPTR_ERROR_MINIPOS_MODE_FR_DISABLED                = 263
	LIBFPTR_ERROR_ENTRY_NOT_FOUND_IN_OTP                  = 264
	LIBFPTR_ERROR_EXCISABLE_COMMODITY_WITHOUT_EXCISE      = 265
	LIBFPTR_ERROR_BARCODE_TYPE_NOT_SUPPORTED              = 266
	LIBFPTR_ERROR_OVERLAY_DATA_OVERFLOW                   = 267
	LIBFPTR_ERROR_INVALID_MODULE_ADDRESS                  = 268
	LIBFPTR_ERROR_ECR_MODEL_NOT_SUPPORTED                 = 269
	LIBFPTR_ERROR_PAID_NOT_REQUIRED                       = 270
	LIBFPTR_ERROR_NON_PRINTABLE_CHAR                      = 271
	TXT_ERROR_INVALID_USER_TAG                            = 272
)

const (
	LIBFPTR_ERROR_BASE_MARKING                        = 400
	LIBFPTR_ERROR_MARKING_CODE_VALIDATION_IN_PROGRESS = 401
	LIBFPTR_ERROR_NO_CONNECTION_WITH_SERVER           = 402
	LIBFPTR_ERROR_MARKING_CODE_VALIDATION_CANCELED    = 403
	LIBFPTR_ERROR_INVALID_MARKING_CODE_STATUS         = 404
	LIBFPTR_ERROR_INVALID_GS1                         = 405
	LIBFPTR_ERROR_MARKING_WORK_DENIED                 = 406
	LIBFPTR_ERROR_MARKING_WORK_TEMPORARY_BLOCKED      = 407
	LIBFPTR_ERROR_MARKS_OVERFLOW                      = 408
	LIBFPTR_ERROR_INVALID_MARKING_CODE                = 409
	LIBFPTR_ERROR_INVALID_STATE                       = 410
	LIBFPTR_ERROR_OFD_EXCHANGE                        = 411
	LIBFPTR_ERROR_INVALID_MEASUREMENT_UNIT            = 412
	LIBFPTR_ERROR_OPERATION_DENIED_IN_CURRENT_FFD     = 413
	LIBFPTR_ERROR_MARKING_OPERATION_DENIED            = 414
	LIBFPTR_ERROR_NO_DATA_TO_SEND                     = 415
	LIBFPTR_ERROR_NO_MARKED_POSITION                  = 416
	LIBFPTR_ERROR_HAS_NOT_SEND_NOTICES                = 417
	LIBFPTR_ERROR_UPDATE_KEYS_REQUIRED                = 418
	LIBFPTR_ERROR_UPDATE_KEYS_SERVICE                 = 419
	LIBFPTR_ERROR_MARK_NOT_CHECKED                    = 420
	LIBFPTR_ERROR_MARK_CHECK_TIMEOUT_EXPIRED          = 421
	LIBFPTR_ERROR_NO_MARKING_CODE_IN_TABLE            = 422
	LIBFPTR_ERROR_CHEKING_MARK_IN_PROGRESS            = 423
	LIBFPTR_ERROR_INVALID_SERVER_ADDRESS              = 424
	LIBFPTR_ERROR_UPDATE_KEYS_TIMEOUT                 = 425
	LIBFPTR_ERROR_PROPERTY_FOR_MARKING_POSITION_ONLY  = 426
)

const (
	LIBFPTR_ERROR_BASE_WEB                       = 500
	LIBFPTR_ERROR_RECEIPT_PARSE_ERROR            = 501
	LIBFPTR_ERROR_INTERRUPTED_BY_PREVIOUS_ERRORS = 502
	LIBFPTR_ERROR_DRIVER_SCRIPT_ERROR            = 503
	LIBFPTR_ERROR_VALIDATE_FUNC_NOT_FOUND        = 504
	LIBFPTR_ERROR_WEB_FAIL                       = 505
)

const (
	LIBFPTR_PORT_COM       = 0
	LIBFPTR_PORT_USB       = 1
	LIBFPTR_PORT_TCPIP     = 2
	LIBFPTR_PORT_BLUETOOTH = 3
)

const (
	LIBFPTR_PORT_BITS_7 = 7
	LIBFPTR_PORT_BITS_8 = 8
)

const (
	LIBFPTR_PORT_PARITY_NO    = 0
	LIBFPTR_PORT_PARITY_ODD   = 1
	LIBFPTR_PORT_PARITY_EVEN  = 2
	LIBFPTR_PORT_PARITY_MARK  = 3
	LIBFPTR_PORT_PARITY_SPACE = 4
)

const (
	LIBFPTR_PORT_SB_1   = 0
	LIBFPTR_PORT_SB_1_5 = 1
	LIBFPTR_PORT_SB_2   = 2
)

const (
	LIBFPTR_BT_EAN_8            = 0
	LIBFPTR_BT_EAN_13           = 1
	LIBFPTR_BT_UPC_A            = 2
	LIBFPTR_BT_UPC_E            = 3
	LIBFPTR_BT_CODE_39          = 4
	LIBFPTR_BT_CODE_93          = 5
	LIBFPTR_BT_CODE_128         = 6
	LIBFPTR_BT_CODABAR          = 7
	LIBFPTR_BT_ITF              = 8
	LIBFPTR_BT_ITF_14           = 9
	LIBFPTR_BT_GS1_128          = 10
	LIBFPTR_BT_QR               = 11
	LIBFPTR_BT_PDF417           = 12
	LIBFPTR_BT_AZTEC            = 13
	LIBFPTR_BT_CODE_39_EXTENDED = 14
)

const (
	LIBFPTR_BC_DEFAULT = 0
	LIBFPTR_BC_0       = 1
	LIBFPTR_BC_1       = 2
	LIBFPTR_BC_2       = 3
	LIBFPTR_BC_3       = 4
	LIBFPTR_BC_4       = 5
	LIBFPTR_BC_5       = 6
	LIBFPTR_BC_6       = 7
	LIBFPTR_BC_7       = 8
	LIBFPTR_BC_8       = 9
)

const (
	LIBFPTR_TM_POSITION = 0
	LIBFPTR_TM_UNIT     = 1
)

const (
	LIBFPTR_SCT_OVERALL = 0
	LIBFPTR_SCT_FORWARD = 1
)

const (
	LIBFPTR_CT_ROLLUP     = 0
	LIBFPTR_CT_RESETTABLE = 1
)

const (
	LIBFPTR_SS_CLOSED  = 0
	LIBFPTR_SS_OPENED  = 1
	LIBFPTR_SS_EXPIRED = 2
)

const (
	LIBFPTR_CT_FULL = 0
	LIBFPTR_CT_PART = 1
)

const (
	LIBFPTR_ALIGNMENT_LEFT   = 0
	LIBFPTR_ALIGNMENT_CENTER = 1
	LIBFPTR_ALIGNMENT_RIGHT  = 2
)

const (
	LIBFPTR_TW_NONE  = 0
	LIBFPTR_TW_WORDS = 1
	LIBFPTR_TW_CHARS = 2
)

const (
	LIBFPTR_FNT_DEBUG   = 0
	LIBFPTR_FNT_RELEASE = 1
	LIBFPTR_FNT_UNKNOWN = 2
)

const (
	LIBFPTR_RT_CLOSE_SHIFT                    = 0
	LIBFPTR_RT_X                              = 1
	LIBFPTR_RT_LAST_DOCUMENT                  = 2
	LIBFPTR_RT_OFD_EXCHANGE_STATUS            = 3
	LIBFPTR_RT_KKT_DEMO                       = 4
	LIBFPTR_RT_KKT_INFO                       = 5
	LIBFPTR_RT_OFD_TEST                       = 6
	LIBFPTR_RT_FN_DOC_BY_NUMBER               = 7
	LIBFPTR_RT_QUANTITY                       = 8
	LIBFPTR_RT_DEPARTMENTS                    = 9
	LIBFPTR_RT_OPERATORS                      = 10
	LIBFPTR_RT_HOURS                          = 11
	LIBFPTR_RT_FN_REGISTRATIONS               = 12
	LIBFPTR_RT_FN_SHIFT_TOTAL_COUNTERS        = 13
	LIBFPTR_RT_FN_TOTAL_COUNTERS              = 14
	LIBFPTR_RT_FN_NOT_SENT_DOCUMENTS_COUNTERS = 15
	LIBFPTR_RT_COMMODITIES_BY_TAXATION_TYPES  = 16
	LIBFPTR_RT_COMMODITIES_BY_DEPARTMENTS     = 17
	LIBFPTR_RT_COMMODITIES_BY_SUMS            = 18
	LIBFPTR_RT_START_SERVICE                  = 19
	LIBFPTR_RT_DISCOUNTS                      = 20
	LIBFPTR_RT_JOURNAL_DOCUMENT_BY_NUMBERS    = 21
	LIBFPTR_RT_JOURNAL_DOCUMENT_BY_SHIFTS     = 22
	LIBFPTR_RT_CLOSE_SHIFT_REPORTS            = 23
)

const (
	LIBFPTR_PT_CASH           = 0
	LIBFPTR_PT_ELECTRONICALLY = 1
	LIBFPTR_PT_PREPAID        = 2
	LIBFPTR_PT_CREDIT         = 3
	LIBFPTR_PT_OTHER          = 4
	LIBFPTR_PT_6              = 5
	LIBFPTR_PT_7              = 6
	LIBFPTR_PT_8              = 7
	LIBFPTR_PT_9              = 8
	LIBFPTR_PT_10             = 9
)

const (
	LIBFPTR_TAX_DEPARTMENT = 0
	LIBFPTR_TAX_VAT18      = 1
	LIBFPTR_TAX_VAT10      = 2
	LIBFPTR_TAX_VAT118     = 3
	LIBFPTR_TAX_VAT110     = 4
	LIBFPTR_TAX_VAT0       = 5
	LIBFPTR_TAX_NO         = 6
	LIBFPTR_TAX_VAT20      = 7
	LIBFPTR_TAX_VAT120     = 8
	LIBFPTR_TAX_INVALID    = 9
)

const (
	LIBFPTR_EXTERNAL_DEVICE_DISPLAY         = 0
	LIBFPTR_EXTERNAL_DEVICE_PINPAD          = 1
	LIBFPTR_EXTERNAL_DEVICE_MODEM           = 2
	LIBFPTR_EXTERNAL_DEVICE_BARCODE_SCANNER = 3
)

const (
	LIBFPTR_DT_STATUS                           = 0
	LIBFPTR_DT_CASH_SUM                         = 1
	LIBFPTR_DT_UNIT_VERSION                     = 2
	LIBFPTR_DT_PICTURE_INFO                     = 3
	LIBFPTR_DT_LICENSE_ACTIVATED                = 4
	LIBFPTR_DT_REGISTRATIONS_SUM                = 5
	LIBFPTR_DT_REGISTRATIONS_COUNT              = 6
	LIBFPTR_DT_PAYMENT_SUM                      = 7
	LIBFPTR_DT_CASHIN_SUM                       = 8
	LIBFPTR_DT_CASHIN_COUNT                     = 9
	LIBFPTR_DT_CASHOUT_SUM                      = 10
	LIBFPTR_DT_CASHOUT_COUNT                    = 11
	LIBFPTR_DT_REVENUE                          = 12
	LIBFPTR_DT_DATE_TIME                        = 13
	LIBFPTR_DT_SHIFT_STATE                      = 14
	LIBFPTR_DT_RECEIPT_STATE                    = 15
	LIBFPTR_DT_SERIAL_NUMBER                    = 16
	LIBFPTR_DT_MODEL_INFO                       = 17
	LIBFPTR_DT_RECEIPT_LINE_LENGTH              = 18
	LIBFPTR_DT_CUTTER_RESOURCE                  = 19
	LIBFPTR_DT_STEP_RESOURCE                    = 20
	LIBFPTR_DT_TERMAL_RESOURCE                  = 21
	LIBFPTR_DT_ENVD_MODE                        = 22
	LIBFPTR_DT_SHIFT_TAX_SUM                    = 23
	LIBFPTR_DT_RECEIPT_TAX_SUM                  = 24
	LIBFPTR_DT_NON_NULLABLE_SUM                 = 25
	LIBFPTR_DT_RECEIPT_COUNT                    = 26
	LIBFPTR_DT_CANCELLATION_COUNT_ALL           = 27
	LIBFPTR_DT_CANCELLATION_SUM                 = 28
	LIBFPTR_DT_CANCELLATION_SUM_ALL             = 29
	LIBFPTR_DT_POWER_SOURCE_STATE               = 30
	LIBFPTR_DT_CANCELLATION_COUNT               = 31
	LIBFPTR_DT_NON_NULLABLE_SUM_BY_PAYMENTS     = 32
	LIBFPTR_DT_PRINTER_TEMPERATURE              = 33
	LIBFPTR_DT_FATAL_STATUS                     = 34
	LIBFPTR_DT_MAC_ADDRESS                      = 35
	LIBFPTR_DT_DEVICE_UPTIME                    = 36
	LIBFPTR_DT_RECEIPT_BYTE_COUNT               = 37
	LIBFPTR_DT_DISCOUNT_AND_SURCHARGE_SUM       = 38
	LIBFPTR_DT_LK_USER_CODE                     = 39
	LIBFPTR_DT_LAST_SENT_OFD_DOCUMENT_DATE_TIME = 40
	LIBFPTR_DT_SHORT_STATUS                     = 41
	LIBFPTR_DT_PICTURES_ARRAY_INFO              = 42
	LIBFPTR_DT_ETHERNET_INFO                    = 43
	LIBFPTR_DT_SCRIPTS_INFO                     = 44
	LIBFPTR_DT_SHIFT_TOTALS                     = 45
	LIBFPTR_DT_WIFI_INFO                        = 46
	LIBFPTR_DT_FONT_INFO                        = 47
	LIBFPTR_DT_SOFTLOCK_STATUS                  = 48
	LIBFPTR_DT_LAST_SENT_ISM_NOTICE_DATE_TIME   = 49
	LIBFPTR_DT_MCU_INFO                         = 50
	LIBFPTR_DT_MODULE_ADDRESS                   = 51
	LIBFPTR_DT_CACHE_REQUISITES                 = 52
	LIBFPTR_DT_DEPARTMENT_SUM                   = 53
	LIBFPTR_DT_MCU_TEMPERATURE                  = 54
	LIBFPTR_DT_AVAILABLE_OPERATIONS             = 55
	LIBFPTR_DT_PATTERN_PARAMETERS               = 56
)

const (
	LIBFPTR_FNDT_TAG_VALUE                = 0
	LIBFPTR_FNDT_OFD_EXCHANGE_STATUS      = 1
	LIBFPTR_FNDT_FN_INFO                  = 2
	LIBFPTR_FNDT_LAST_REGISTRATION        = 3
	LIBFPTR_FNDT_LAST_RECEIPT             = 4
	LIBFPTR_FNDT_LAST_DOCUMENT            = 5
	LIBFPTR_FNDT_SHIFT                    = 6
	LIBFPTR_FNDT_FFD_VERSIONS             = 7
	LIBFPTR_FNDT_VALIDITY                 = 8
	LIBFPTR_FNDT_REG_INFO                 = 9
	LIBFPTR_FNDT_DOCUMENTS_COUNT_IN_SHIFT = 10
	LIBFPTR_FNDT_ERRORS                   = 11
	LIBFPTR_FNDT_TICKET_BY_DOC_NUMBER     = 12
	LIBFPTR_FNDT_DOCUMENT_BY_NUMBER       = 13
	LIBFPTR_FNDT_REGISTRATION_TLV         = 14
	LIBFPTR_FNDT_ERROR_DETAIL             = 15
	LIBFPTR_FNDT_VALIDITY_DAYS            = 16
	LIBFPTR_FNDT_FREE_MEMORY              = 17
	LIBFPTR_FNDT_TOTALS                   = 18
	LIBFPTR_FNDT_ISM_ERRORS               = 19
	LIBFPTR_FNDT_ISM_EXCHANGE_STATUS      = 20
	LIBFPTR_FNDT_MARKING_MODE_STATUS      = 21
	LIBFPTR_FNDT_CHECK_MARK_TIME          = 22
	LIBFPTR_FNDT_RECEIPT_SIZE             = 23
)

const (
	LIBFPTR_UT_FIRMWARE      = 0
	LIBFPTR_UT_CONFIGURATION = 1
	LIBFPTR_UT_TEMPLATES     = 2
	LIBFPTR_UT_CONTROL_UNIT  = 3
	LIBFPTR_UT_BOOT          = 4
)

const (
	LIBFPTR_FNOP_REGISTRATION      = 0
	LIBFPTR_FNOP_CHANGE_FN         = 1
	LIBFPTR_FNOP_CHANGE_PARAMETERS = 2
	LIBFPTR_FNOP_CLOSE_ARCHIVE     = 3
)

const (
	LIBFPTR_OFD_CHANNEL_NONE  = 0
	LIBFPTR_OFD_CHANNEL_USB   = 1
	LIBFPTR_OFD_CHANNEL_PROTO = 2
)

const (
	LIBFPTR_PST_POWER_SUPPLY = 0
	LIBFPTR_PST_RTC_BATTERY  = 1
	LIBFPTR_PST_BATTERY      = 2
)

const (
	LIBFPTR_RT_LAST_DOCUMENT_LINES     = 0
	LIBFPTR_RT_FN_DOCUMENT_TLVS        = 1
	LIBFPTR_RT_EXEC_USER_SCRIPT        = 2
	LIBFPTR_RT_FIRMWARE                = 3
	LIBFPTR_RT_LICENSES                = 4
	LIBFPTR_RT_FN_REGISTRATION_TLVS    = 5
	LIBFPTR_RT_PARSE_COMPLEX_ATTR      = 6
	LIBFPTR_RT_FN_SUM_COUNTERS         = 7
	LIBFPTR_RT_FN_QUANTITY_COUNTERS    = 8
	LIBFPTR_RT_FN_UNSENT_DOCS_COUNTERS = 9
	LIBFPTR_RT_SETTINGS                = 10
	LIBFPTR_RT_RUN_COMMAND             = 11
)

const (
	LIBFPTR_LOG_ERROR = 0
	LIBFPTR_LOG_WARN  = 1
	LIBFPTR_LOG_INFO  = 2
	LIBFPTR_LOG_DEBUG = 3
)

const (
	LIBFPTR_NT_FURS      = 0
	LIBFPTR_NT_MEDICINES = 1
	LIBFPTR_NT_TOBACCO   = 2
	LIBFPTR_NT_SHOES     = 3
)

const (
	LIBFPTR_UMO_GET_SIZE     = 0
	LIBFPTR_UMO_READ_DATA    = 1
	LIBFPTR_UMO_WRITE_DATA   = 2
	LIBFPTR_UMO_READ_STRING  = 3
	LIBFPTR_UMO_WRITE_STRING = 4
	LIBFPTR_UMO_COMMIT       = 5
)

const (
	LIBFPTR_GUI_PARENT_NATIVE = 0
	LIBFPTR_GUI_PARENT_QT     = 1
)

const (
	LIBFPTR_DEFER_NONE    = 0
	LIBFPTR_DEFER_PRE     = 1
	LIBFPTR_DEFER_POST    = 2
	LIBFPTR_DEFER_OVERLAY = 3
)

const (
	LIBFPTR_TAG_TYPE_STLV      = 0
	LIBFPTR_TAG_TYPE_STRING    = 1
	LIBFPTR_TAG_TYPE_ARRAY     = 2
	LIBFPTR_TAG_TYPE_FVLN      = 3
	LIBFPTR_TAG_TYPE_BITS      = 4
	LIBFPTR_TAG_TYPE_BYTE      = 5
	LIBFPTR_TAG_TYPE_VLN       = 6
	LIBFPTR_TAG_TYPE_UINT_16   = 7
	LIBFPTR_TAG_TYPE_UINT_32   = 8
	LIBFPTR_TAG_TYPE_UNIX_TIME = 9
	LIBFPTR_TAG_TYPE_BOOL      = 10
)

const (
	LIBFPTR_FT_BYTE_ARRAY       = 0
	LIBFPTR_FT_BIN              = 1
	LIBFPTR_FT_BCD              = 2
	LIBFPTR_FT_STRING           = 3
	LIBFPTR_FT_STRING_NULL_TERM = 4
)

const (
	LIBFPTR_ST_NUMBER = 0
	LIBFPTR_ST_STRING = 1
	LIBFPTR_ST_BOOL   = 2
)

const (
	LIBFPTR_SCRIPT_EXECUTABLE = 0
	LIBFPTR_SCRIPT_JSON       = 1
	LIBFPTR_SCRIPT_SETTINGS   = 2
	LIBFPTR_SCRIPT_LIBRARY    = 3
)

const (
	LIBFPTR_UCL_UNUSED         = 0
	LIBFPTR_UCL_RECEIPT_TYPE   = 1
	LIBFPTR_UCL_TAXATION_TYPE  = 2
	LIBFPTR_UCL_TAX_TYPE       = 3
	LIBFPTR_UCL_PRODUCT_TYPE   = 4
	LIBFPTR_UCL_PAYMENT_METHOD = 5
	LIBFPTR_UCL_USER_3         = 6
	LIBFPTR_UCL_USER_4         = 7
	LIBFPTR_UCL_USER_5         = 8
	LIBFPTR_UCL_USER_6         = 9
)

const (
	LIBFPTR_FNCT_SHIFT        = 0
	LIBFPTR_FNCT_NON_NULLABLE = 1
)

const (
	LIBFPTR_MCT_OTHER    = 0
	LIBFPTR_MCT_EGAIS_20 = 1
	LIBFPTR_MCT_EGAIS_30 = 2
)

const (
	LIBFPTR_MCT12_UNKNOWN     = 0
	LIBFPTR_MCT12_SHORT       = 1
	LIBFPTR_MCT12_88_CHECK    = 2
	LIBFPTR_MCT12_44_NO_CHECK = 3
	LIBFPTR_MCT12_44_CHECK    = 4
	LIBFPTR_MCT12_4_NO_CHECK  = 5
)

const (
	LIBFPTR_MES_PIECE_SOLD   = 1
	LIBFPTR_MES_DRY_FOR_SALE = 2
	LIBFPTR_MES_PIECE_RETURN = 3
	LIBFPTR_MES_DRY_RETURN   = 4
)

const (
	LIBFPTR_MCS_BLOCK                = 0
	LIBFPTR_MCS_NO_MARK_FOR_CHECK    = 1
	LIBFPTR_MCS_MARK_RECEIVE_B1      = 2
	LIBFPTR_MCS_MARK_STATE_QUERY_B5  = 3
	LIBFPTR_MCS_MARK_STATE_ANSWER_B6 = 4
)

const (
	LIBFPTR_NFM_LESS_50_PERCENT       = 0
	LIBFPTR_NFM_FROM_50_TO_80_PERCENT = 1
	LIBFPTR_NFM_FROM_80_TO_90_PERCENT = 2
	LIBFPTR_NFM_MORE_90_PERCENT       = 3
	LIBFPTR_NFM_OUT_OF_MEMORY         = 4
)

const (
	LIBFPTR_OIS_ESTIMATED_STATUS_CORRECT   = 1
	LIBFPTR_OIS_ESTIMATED_STATUS_INCORRECT = 2
	LIBFPTR_OIS_SALE_STOPPED               = 3
)

const (
	LIBFPTR_ORR_CORRECT      = 0
	LIBFPTR_ORR_INCORRECT    = 1
	LIBFPTR_ORR_UNRECOGNIZED = 2
)

const (
	LIBFPTR_CER_CHECKED        = 0
	LIBFPTR_CER_TYPE_INCORRECT = 1
	LIBFPTR_CER_NO_KEYS        = 2
	LIBFPTR_CER_NO_GS1         = 3
	LIBFPTR_CER_OTHER          = 4
)

const (
	LIBFPTR_MCS_NOT_EXECUTED       = 0
	LIBFPTR_MCS_EXECUTED           = 1
	LIBFPTR_MCS_IS_OVER            = 2
	LIBFPTR_MCS_RESULT_IS_RECIEVED = 3
)

const (
	LIBFPTR_MCT_AUTONOMOUS      = 0
	LIBFPTR_MCT_WAIT_FOR_RESULT = 1
	LIBFPTR_MCT_RESULT_NOT_WAIT = 2
	LIBFPTR_MCT_QUERY_NOT_SEND  = 3
)

const (
	LIBFPTR_MCST_WAITING_FOR_TASK   = 0
	LIBFPTR_MCST_OPENING_CONNECTION = 1
	LIBFPTR_MCST_SENDING            = 2
	LIBFPTR_MCST_WAITING_FOR_RESULT = 3
	LIBFPTR_MCST_GETTING_RESULT     = 4
	LIBFPTR_MCST_DECODE_RESULT      = 5
	LIBFPTR_MCST_TASK_IS_OVER       = 6
	LIBFPTR_MCST_WAITING_FOR_REPEAT = 7
)

const (
	LIBFPTR_SILENT_REBOOT_NO                  = 0
	LIBFPTR_SILENT_REBOOT_AFTER_SESSION_CLOSE = 1
	LIBFPTR_SILENT_REBOOT_BEFORE_SESSION_OPEN = 2
)

const (
	LIBFPTR_ERROR_BASE_RPC            = 600
	LIBFPTR_ERROR_RCP_SERVER_BUSY     = 601
	LIBFPTR_ERROR_RCP_SERVER_VERSION  = 602
	LIBFPTR_ERROR_RCP_SERVER_EXCHANGE = 603
)

const (
	LIBFPTR_ERROR_MARKING_END                        = 499
	LIBFPTR_ERROR_WEB_END                            = 599
	LIBFPTR_OFD_CHANNEL_AUTO                         = 2
	LIBFPTR_SETTING_LIBRARY_PATH                     = "LibraryPath"
	LIBFPTR_SETTING_MODEL                            = "Model"
	LIBFPTR_SETTING_PORT                             = "Port"
	LIBFPTR_SETTING_BAUDRATE                         = "BaudRate"
	LIBFPTR_SETTING_BITS                             = "Bits"
	LIBFPTR_SETTING_PARITY                           = "Parity"
	LIBFPTR_SETTING_STOPBITS                         = "StopBits"
	LIBFPTR_SETTING_IPADDRESS                        = "IPAddress"
	LIBFPTR_SETTING_IPPORT                           = "IPPort"
	LIBFPTR_SETTING_MACADDRESS                       = "MACAddress"
	LIBFPTR_SETTING_COM_FILE                         = "ComFile"
	LIBFPTR_SETTING_USB_DEVICE_PATH                  = "UsbDevicePath"
	LIBFPTR_SETTING_BT_AUTOENABLE                    = "AutoEnableBluetooth"
	LIBFPTR_SETTING_BT_AUTODISABLE                   = "AutoDisableBluetooth"
	LIBFPTR_SETTING_ACCESS_PASSWORD                  = "AccessPassword"
	LIBFPTR_SETTING_USER_PASSWORD                    = "UserPassword"
	LIBFPTR_SETTING_OFD_CHANNEL                      = "OfdChannel"
	LIBFPTR_SETTING_EXISTED_COM_FILES                = "ExistedComFiles"
	LIBFPTR_SETTING_SCRIPTS_PATH                     = "ScriptsPath"
	LIBFPTR_SETTING_DOCUMENTS_JOURNAL_PATH           = "DocumentsJournalPath"
	LIBFPTR_SETTING_USE_DOCUMENTS_JOURNAL            = "UseDocumentsJournal"
	LIBFPTR_SETTING_AUTO_RECONNECT                   = "AutoReconnect"
	LIBFPTR_SETTING_INVERT_CASH_DRAWER_STATUS        = "InvertCashDrawerStatus"
	LIBFPTR_SETTING_REMOTE_SERVER_ADDR               = "RemoteServerAddr"
	LIBFPTR_SETTING_REMOTE_SERVER_CONNECTION_TIMEOUT = "RemoteServerConnectionTimeout"
	LIBFPTR_SETTING_VALIDATE_MARK_WITH_FNM_ONLY      = "ValidateMarksWithFnmOnly"
	LIBFPTR_SETTING_AUTO_MEASUREMENT_UNIT            = "AutoMeasurementUnit"
	LIBFPTR_SETTING_SILENT_REBOOT                    = "SilentReboot"
	LIBFPTR_MODEL_UNKNOWN                            = 0
	LIBFPTR_MODEL_ATOL_25F                           = 57
	LIBFPTR_MODEL_ATOL_30F                           = 61
	LIBFPTR_MODEL_ATOL_55F                           = 62
	LIBFPTR_MODEL_ATOL_22F                           = 63
	LIBFPTR_MODEL_ATOL_52F                           = 64
	LIBFPTR_MODEL_ATOL_11F                           = 67
	LIBFPTR_MODEL_ATOL_77F                           = 69
	LIBFPTR_MODEL_ATOL_90F                           = 72
	LIBFPTR_MODEL_ATOL_60F                           = 75
	LIBFPTR_MODEL_ATOL_42FS                          = 77
	LIBFPTR_MODEL_ATOL_15F                           = 78
	LIBFPTR_MODEL_ATOL_50F                           = 80
	LIBFPTR_MODEL_ATOL_20F                           = 81
	LIBFPTR_MODEL_ATOL_91F                           = 82
	LIBFPTR_MODEL_ATOL_92F                           = 84
	LIBFPTR_MODEL_ATOL_SIGMA_10                      = 86
	LIBFPTR_MODEL_ATOL_27F                           = 87
	LIBFPTR_MODEL_ATOL_SIGMA_7F                      = 90
	LIBFPTR_MODEL_ATOL_SIGMA_8F                      = 91
	LIBFPTR_MODEL_ATOL_1F                            = 93
	LIBFPTR_MODEL_KAZNACHEY_FA                       = 76
	LIBFPTR_MODEL_ATOL_22V2F                         = 95
	LIBFPTR_MODEL_ATOL_AUTO                          = 500
	LIBFPTR_MODEL_ATOL_47FA                          = 48
	LIBFPTR_MODEL_ATOL_PT_5F                         = 89
	LIBFPTR_MODEL_ATOL_42FA                          = 70
	LIBFPTR_MODEL_ALLIANCE_20F                       = 50
	LIBFPTR_PORT_BR_1200                             = 1200
	LIBFPTR_PORT_BR_2400                             = 2400
	LIBFPTR_PORT_BR_4800                             = 4800
	LIBFPTR_PORT_BR_9600                             = 9600
	LIBFPTR_PORT_BR_19200                            = 19200
	LIBFPTR_PORT_BR_38400                            = 38400
	LIBFPTR_PORT_BR_57600                            = 57600
	LIBFPTR_PORT_BR_115200                           = 115200
	LIBFPTR_PORT_BR_230400                           = 230400
	LIBFPTR_PORT_BR_460800                           = 460800
	LIBFPTR_PORT_BR_921600                           = 921600
	LIBFPTR_FNS_INITIAL                              = 0
	LIBFPTR_FNS_CONFIGURED                           = 1
	LIBFPTR_FNS_FISCAL_MODE                          = 3
	LIBFPTR_FNS_POSTFISCAL_MODE                      = 7
	LIBFPTR_FNS_ACCESS_ARCHIVE                       = 15
	LIBFPTR_RT_CLOSED                                = 0
	LIBFPTR_RT_SELL                                  = 1
	LIBFPTR_RT_SELL_RETURN                           = 2
	LIBFPTR_RT_SELL_CORRECTION                       = 7
	LIBFPTR_RT_SELL_RETURN_CORRECTION                = 8
	LIBFPTR_RT_BUY                                   = 4
	LIBFPTR_RT_BUY_RETURN                            = 5
	LIBFPTR_RT_BUY_CORRECTION                        = 9
	LIBFPTR_RT_BUY_RETURN_CORRECTION                 = 10
	LIBFPTR_FFD_UNKNOWN                              = 0
	LIBFPTR_FFD_1_0                                  = 100
	LIBFPTR_FFD_1_0_5                                = 105
	LIBFPTR_FFD_1_1                                  = 110
	LIBFPTR_FFD_1_2                                  = 120
	LIBFPTR_TT_DEFAULT                               = 0
	LIBFPTR_TT_OSN                                   = 1
	LIBFPTR_TT_USN_INCOME                            = 2
	LIBFPTR_TT_USN_INCOME_OUTCOME                    = 4
	LIBFPTR_TT_ENVD                                  = 8
	LIBFPTR_TT_ESN                                   = 16
	LIBFPTR_TT_PATENT                                = 32
	LIBFPTR_AT_NONE                                  = 0
	LIBFPTR_AT_BANK_PAYING_AGENT                     = 1
	LIBFPTR_AT_BANK_PAYING_SUBAGENT                  = 2
	LIBFPTR_AT_PAYING_AGENT                          = 4
	LIBFPTR_AT_PAYING_SUBAGENT                       = 8
	LIBFPTR_AT_ATTORNEY                              = 16
	LIBFPTR_AT_COMMISSION_AGENT                      = 32
	LIBFPTR_AT_ANOTHER                               = 64
	LIBFPTR_DT_CLOSED                                = 0
	LIBFPTR_DT_RECEIPT_SELL                          = 1
	LIBFPTR_DT_RECEIPT_SELL_RETURN                   = 2
	LIBFPTR_DT_RECEIPT_BUY                           = 3
	LIBFPTR_DT_RECEIPT_BUY_RETURN                    = 4
	LIBFPTR_DT_OPEN_SHIFT                            = 5
	LIBFPTR_DT_CLOSE_SHIFT                           = 6
	LIBFPTR_DT_REGISTRATION                          = 7
	LIBFPTR_DT_CLOSE_ARCHIVE                         = 8
	LIBFPTR_DT_OFD_EXCHANGE_STATUS                   = 11
	LIBFPTR_DT_RECEIPT_SELL_CORRECTION               = 12
	LIBFPTR_DT_RECEIPT_SELL_RETURN_CORRECTION        = 13
	LIBFPTR_DT_RECEIPT_BUY_CORRECTION                = 14
	LIBFPTR_DT_RECEIPT_BUY_RETURN_CORRECTION         = 15
	LIBFPTR_DT_DOCUMENT_SERVICE                      = 20
	LIBFPTR_DT_DOCUMENT_COPY                         = 21
	LIBFPTR_FN_DOC_REGISTRATION                      = 1
	LIBFPTR_FN_DOC_OPEN_SHIFT                        = 2
	LIBFPTR_FN_DOC_RECEIPT                           = 3
	LIBFPTR_FN_DOC_BSO                               = 4
	LIBFPTR_FN_DOC_CLOSE_SHIFT                       = 5
	LIBFPTR_FN_DOC_CLOSE_FN                          = 6
	LIBFPTR_FN_DOC_OPERATOR_CONFIRMATION             = 7
	LIBFPTR_FN_DOC_REREGISTRATION                    = 11
	LIBFPTR_FN_DOC_EXCHANGE_STATUS                   = 21
	LIBFPTR_FN_DOC_CORRECTION                        = 31
	LIBFPTR_FN_DOC_BSO_CORRECTION                    = 41
	LIBFPTR_FWT_FIRMWARE                             = 0
	LIBFPTR_FWT_SCRIPTS                              = 2
	LIBFPTR_UCF_CALC_SUMS                            = 1
	LIBFPTR_UCF_CALC_QUANTITIES                      = 2
	LIBFPTR_UCF_CALC_SUMS_OTHERS                     = 4
	LIBFPTR_UCF_CALC_QUANTITIES_OTHERS               = 8
	LIBFPTR_UC_OTHERS                                = 4294967295
	LIBFPTR_MCT12_AUTO                               = 256
	LIBFPTR_MES_UNCHANGED                            = 255
	LIBFPTR_IU_PIECE                                 = 0
	LIBFPTR_IU_GRAM                                  = 10
	LIBFPTR_IU_KILOGRAM                              = 11
	LIBFPTR_IU_TON                                   = 12
	LIBFPTR_IU_CENTIMETER                            = 20
	LIBFPTR_IU_DECIMETER                             = 21
	LIBFPTR_IU_METER                                 = 22
	LIBFPTR_IU_SQUARE_CENTIMETER                     = 30
	LIBFPTR_IU_SQUARE_DECIMETER                      = 31
	LIBFPTR_IU_SQUARE_METER                          = 32
	LIBFPTR_IU_MILLILITER                            = 40
	LIBFPTR_IU_LITER                                 = 41
	LIBFPTR_IU_CUBIC_METER                           = 42
	LIBFPTR_IU_KILOWATT_HOUR                         = 50
	LIBFPTR_IU_GKAL                                  = 51
	LIBFPTR_IU_DAY                                   = 70
	LIBFPTR_IU_HOUR                                  = 71
	LIBFPTR_IU_MINUTE                                = 72
	LIBFPTR_IU_SECOND                                = 73
	LIBFPTR_IU_KILOBYTE                              = 80
	LIBFPTR_IU_MEGABYTE                              = 81
	LIBFPTR_IU_GIGABYTE                              = 82
	LIBFPTR_IU_TERABYTE                              = 83
	LIBFPTR_IU_OTHER                                 = 255
	LIBFPTR_ERROR_USERS_SCRIPTS_BASE                 = 1000
	LIBFPTR_PLATFORM_UNKNOWN                         = 0
	LIBFPTR_PLATFORM_25                              = 25
	LIBFPTR_PLATFORM_50                              = 50
	LIBFPTR_ERROR_USERS_SCRIPTS_END                  = 1999
	LIBFPTR_ERROR_RPC_END                            = 699
)

type IFptr struct {
	nativePointer C.libfptr_handle
	functions     *functionPointers
}

type functionPointers struct {
	libfptr_create                            C.libfptr_create_func
	libfptr_create_with_id                    C.libfptr_create_with_id_func
	libfptr_destroy                           C.libfptr_destroy_func
	libfptr_get_version_string                C.libfptr_get_version_string_func
	libfptr_set_settings                      C.libfptr_set_settings_func
	libfptr_get_settings                      C.libfptr_get_settings_func
	libfptr_set_single_setting                C.libfptr_set_single_setting_func
	libfptr_get_single_setting                C.libfptr_get_single_setting_func
	libfptr_is_opened                         C.libfptr_is_opened_func
	libfptr_error_code                        C.libfptr_error_code_func
	libfptr_error_description                 C.libfptr_error_description_func
	libfptr_error_recommendation              C.libfptr_error_recommendation_func
	libfptr_reset_error                       C.libfptr_reset_error_func
	libfptr_set_param_bool                    C.libfptr_set_param_bool_func
	libfptr_set_param_int                     C.libfptr_set_param_int_func
	libfptr_set_param_double                  C.libfptr_set_param_double_func
	libfptr_set_param_str                     C.libfptr_set_param_str_func
	libfptr_set_param_bytearray               C.libfptr_set_param_bytearray_func
	libfptr_set_param_datetime                C.libfptr_set_param_datetime_func
	libfptr_set_non_printable_param_bool      C.libfptr_set_param_bool_func
	libfptr_set_non_printable_param_int       C.libfptr_set_param_int_func
	libfptr_set_non_printable_param_double    C.libfptr_set_param_double_func
	libfptr_set_non_printable_param_str       C.libfptr_set_param_str_func
	libfptr_set_non_printable_param_bytearray C.libfptr_set_param_bytearray_func
	libfptr_set_non_printable_param_datetime  C.libfptr_set_param_datetime_func
	libfptr_get_param_bool                    C.libfptr_get_param_bool_func
	libfptr_get_param_int                     C.libfptr_get_param_int_func
	libfptr_get_param_double                  C.libfptr_get_param_double_func
	libfptr_get_param_str                     C.libfptr_get_param_str_func
	libfptr_get_param_bytearray               C.libfptr_get_param_bytearray_func
	libfptr_get_param_datetime                C.libfptr_get_param_datetime_func
	libfptr_is_param_available                C.libfptr_is_param_available_func
	libfptr_log_write                         C.libfptr_log_write_func
	libfptr_change_label                      C.libfptr_change_label_func
	libfptr_show_properties                   C.libfptr_show_properties_func
	libfptr_set_user_param_bool               C.libfptr_set_param_bool_func
	libfptr_set_user_param_int                C.libfptr_set_param_int_func
	libfptr_set_user_param_double             C.libfptr_set_param_double_func
	libfptr_set_user_param_str                C.libfptr_set_param_str_func
	libfptr_set_user_param_bytearray          C.libfptr_set_param_bytearray_func
	libfptr_set_user_param_datetime           C.libfptr_set_param_datetime_func

	libfptr_apply_single_settings                C.libfptr_simple_call_func
	libfptr_open                                 C.libfptr_simple_call_func
	libfptr_close                                C.libfptr_simple_call_func
	libfptr_reset_params                         C.libfptr_simple_call_func
	libfptr_run_command                          C.libfptr_simple_call_func
	libfptr_beep                                 C.libfptr_simple_call_func
	libfptr_open_drawer                          C.libfptr_simple_call_func
	libfptr_cut                                  C.libfptr_simple_call_func
	libfptr_device_poweroff                      C.libfptr_simple_call_func
	libfptr_device_reboot                        C.libfptr_simple_call_func
	libfptr_open_shift                           C.libfptr_simple_call_func
	libfptr_reset_summary                        C.libfptr_simple_call_func
	libfptr_init_device                          C.libfptr_simple_call_func
	libfptr_query_data                           C.libfptr_simple_call_func
	libfptr_cash_income                          C.libfptr_simple_call_func
	libfptr_cash_outcome                         C.libfptr_simple_call_func
	libfptr_open_receipt                         C.libfptr_simple_call_func
	libfptr_cancel_receipt                       C.libfptr_simple_call_func
	libfptr_close_receipt                        C.libfptr_simple_call_func
	libfptr_check_document_closed                C.libfptr_simple_call_func
	libfptr_receipt_total                        C.libfptr_simple_call_func
	libfptr_receipt_tax                          C.libfptr_simple_call_func
	libfptr_registration                         C.libfptr_simple_call_func
	libfptr_payment                              C.libfptr_simple_call_func
	libfptr_report                               C.libfptr_simple_call_func
	libfptr_print_text                           C.libfptr_simple_call_func
	libfptr_print_cliche                         C.libfptr_simple_call_func
	libfptr_begin_nonfiscal_document             C.libfptr_simple_call_func
	libfptr_end_nonfiscal_document               C.libfptr_simple_call_func
	libfptr_print_barcode                        C.libfptr_simple_call_func
	libfptr_print_picture                        C.libfptr_simple_call_func
	libfptr_print_picture_by_number              C.libfptr_simple_call_func
	libfptr_upload_picture_from_file             C.libfptr_simple_call_func
	libfptr_clear_pictures                       C.libfptr_simple_call_func
	libfptr_write_device_setting_raw             C.libfptr_simple_call_func
	libfptr_read_device_setting_raw              C.libfptr_simple_call_func
	libfptr_commit_settings                      C.libfptr_simple_call_func
	libfptr_init_settings                        C.libfptr_simple_call_func
	libfptr_reset_settings                       C.libfptr_simple_call_func
	libfptr_write_date_time                      C.libfptr_simple_call_func
	libfptr_write_license                        C.libfptr_simple_call_func
	libfptr_fn_operation                         C.libfptr_simple_call_func
	libfptr_fn_query_data                        C.libfptr_simple_call_func
	libfptr_fn_write_attributes                  C.libfptr_simple_call_func
	libfptr_external_device_power_on             C.libfptr_simple_call_func
	libfptr_external_device_power_off            C.libfptr_simple_call_func
	libfptr_external_device_write_data           C.libfptr_simple_call_func
	libfptr_external_device_read_data            C.libfptr_simple_call_func
	libfptr_operator_login                       C.libfptr_simple_call_func
	libfptr_process_json                         C.libfptr_simple_call_func
	libfptr_read_device_setting                  C.libfptr_simple_call_func
	libfptr_write_device_setting                 C.libfptr_simple_call_func
	libfptr_begin_read_records                   C.libfptr_simple_call_func
	libfptr_read_next_record                     C.libfptr_simple_call_func
	libfptr_end_read_records                     C.libfptr_simple_call_func
	libfptr_user_memory_operation                C.libfptr_simple_call_func
	libfptr_continue_print                       C.libfptr_simple_call_func
	libfptr_init_mgm                             C.libfptr_simple_call_func
	libfptr_util_form_tlv                        C.libfptr_simple_call_func
	libfptr_util_form_nomenclature               C.libfptr_simple_call_func
	libfptr_util_mapping                         C.libfptr_simple_call_func
	libfptr_read_model_flags                     C.libfptr_simple_call_func
	libfptr_line_feed                            C.libfptr_simple_call_func
	libfptr_flash_firmware                       C.libfptr_simple_call_func
	libfptr_soft_lock_init                       C.libfptr_simple_call_func
	libfptr_soft_lock_query_session_code         C.libfptr_simple_call_func
	libfptr_soft_lock_validate                   C.libfptr_simple_call_func
	libfptr_util_calc_tax                        C.libfptr_simple_call_func
	libfptr_download_picture                     C.libfptr_simple_call_func
	libfptr_bluetooth_remove_paired_devices      C.libfptr_simple_call_func
	libfptr_util_tag_info                        C.libfptr_simple_call_func
	libfptr_util_container_versions              C.libfptr_simple_call_func
	libfptr_activate_licenses                    C.libfptr_simple_call_func
	libfptr_remove_licenses                      C.libfptr_simple_call_func
	libfptr_enter_keys                           C.libfptr_simple_call_func
	libfptr_validate_keys                        C.libfptr_simple_call_func
	libfptr_enter_serial_number                  C.libfptr_simple_call_func
	libfptr_get_serial_number_request            C.libfptr_simple_call_func
	libfptr_upload_pixel_buffer                  C.libfptr_simple_call_func
	libfptr_download_pixel_buffer                C.libfptr_simple_call_func
	libfptr_print_pixel_buffer                   C.libfptr_simple_call_func
	libfptr_util_convert_tag_value               C.libfptr_simple_call_func
	libfptr_parse_marking_code                   C.libfptr_simple_call_func
	libfptr_call_script                          C.libfptr_simple_call_func
	libfptr_set_header_lines                     C.libfptr_simple_call_func
	libfptr_set_footer_lines                     C.libfptr_simple_call_func
	libfptr_upload_picture_cliche                C.libfptr_simple_call_func
	libfptr_upload_picture_memory                C.libfptr_simple_call_func
	libfptr_upload_pixel_buffer_cliche           C.libfptr_simple_call_func
	libfptr_upload_pixel_buffer_memory           C.libfptr_simple_call_func
	libfptr_exec_driver_script                   C.libfptr_simple_call_func
	libfptr_upload_driver_script                 C.libfptr_simple_call_func
	libfptr_exec_driver_script_by_id             C.libfptr_simple_call_func
	libfptr_write_universal_counters_settings    C.libfptr_simple_call_func
	libfptr_read_universal_counters_settings     C.libfptr_simple_call_func
	libfptr_query_universal_counters_state       C.libfptr_simple_call_func
	libfptr_reset_universal_counters             C.libfptr_simple_call_func
	libfptr_cache_universal_counters             C.libfptr_simple_call_func
	libfptr_read_universal_counter_sum           C.libfptr_simple_call_func
	libfptr_read_universal_counter_quantity      C.libfptr_simple_call_func
	libfptr_clear_universal_counters_cache       C.libfptr_simple_call_func
	libfptr_disable_ofd_channel                  C.libfptr_simple_call_func
	libfptr_enable_ofd_channel                   C.libfptr_simple_call_func
	libfptr_validate_json                        C.libfptr_simple_call_func
	libfptr_reflection_call                      C.libfptr_simple_call_func
	libfptr_get_remote_server_info               C.libfptr_simple_call_func
	libfptr_begin_marking_code_validation        C.libfptr_simple_call_func
	libfptr_cancel_marking_code_validation       C.libfptr_simple_call_func
	libfptr_get_marking_code_validation_status   C.libfptr_simple_call_func
	libfptr_accept_marking_code                  C.libfptr_simple_call_func
	libfptr_decline_marking_code                 C.libfptr_simple_call_func
	libfptr_update_fnm_keys                      C.libfptr_simple_call_func
	libfptr_write_sales_notice                   C.libfptr_simple_call_func
	libfptr_check_marking_code_validations_ready C.libfptr_simple_call_func
	libfptr_clear_marking_code_validation_result C.libfptr_simple_call_func
	libfptr_ping_marking_server                  C.libfptr_simple_call_func
	libfptr_get_marking_server_status            C.libfptr_simple_call_func
	libfptr_is_driver_locked                     C.libfptr_simple_call_func
	libfptr_get_last_document_journal            C.libfptr_simple_call_func
	libfptr_find_document_in_journal             C.libfptr_simple_call_func
}

func New() *IFptr {
	fptr := IFptr{}
	if err := fptr.create(""); err != nil {
		return nil
	}
	return &fptr
}

func NewSafe() (*IFptr, error) {
	fptr := IFptr{}
	if err := fptr.create(""); err != nil {
		return nil, err
	}
	return &fptr, nil
}

func NewWithPath(path string) (*IFptr, error) {
	fptr := IFptr{}
	if err := fptr.create(path); err != nil {
		return nil, err
	}
	return &fptr, nil
}

func NewWithID(id string) (*IFptr, error) {
	fptr := IFptr{}
	if err := fptr.createWithID(id, ""); err != nil {
		return nil, err
	}
	return &fptr, nil
}

func NewWithIDAndPath(id string, path string) (*IFptr, error) {
	fptr := IFptr{}
	if err := fptr.createWithID(id, path); err != nil {
		return nil, err
	}
	return &fptr, nil
}

func (fptr *IFptr) create(path string) (err error) {
	functions, err := loadLibrary(path)
	if err != nil {
		return
	}
	fptr.nativePointer = C.libfptr_handle(nil)

	r := C.bridge_libfptr_create_func(functions.libfptr_create, &fptr.nativePointer)
	switch r {
	case 0:
	case -2:
		err = errors.New("Invalid [id] format")
		return
	default:
		err = errors.New("Can`t create driver handle")
		return
	}

	fptr.functions = functions
	return nil
}

func (fptr *IFptr) createWithID(id string, path string) (err error) {
	functions, err := loadLibrary(path)
	if err != nil {
		return
	}
	fptr.nativePointer = C.libfptr_handle(nil)
	strID, _ := StringToWcharT(id)

	r := C.bridge_libfptr_create_with_id_func(functions.libfptr_create_with_id, &fptr.nativePointer, strID)
	switch r {
	case 0:
	case -2:
		err = errors.New("Invalid [id] format")
		return
	default:
		err = errors.New("Can`t create driver handle")
		return
	}

	fptr.functions = functions
	return nil
}

func (fptr *IFptr) Destroy() {
	C.bridge_libfptr_destroy_func(fptr.functions.libfptr_destroy, &fptr.nativePointer)
}

func (fptr *IFptr) Version() string {
	return C.GoString(C.bridge_libfptr_get_version_string_func(fptr.functions.libfptr_get_version_string, fptr.nativePointer))
}

func (fptr *IFptr) SetSettings(settings string) error {
	str, _ := StringToWcharT(settings)
	err := C.bridge_libfptr_set_settings_func(fptr.functions.libfptr_set_settings, fptr.nativePointer, str)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) GetSettings() string {
	buf := make([]C.wchar_t, 128)
	size := int(C.bridge_libfptr_get_settings_func(fptr.functions.libfptr_get_settings, fptr.nativePointer, &buf[0], C.int(len(buf))))
	if size > len(buf) {
		buf = make([]C.wchar_t, size)
		C.bridge_libfptr_get_settings_func(fptr.functions.libfptr_get_settings, fptr.nativePointer, &buf[0], C.int(len(buf)))
	}
	str, _ := WcharTToString(&buf[0])
	return str
}

func (fptr *IFptr) SetSingleSetting(key string, value string) {
	k, _ := StringToWcharT(key)
	val, _ := StringToWcharT(value)
	C.bridge_libfptr_set_single_setting_func(fptr.functions.libfptr_set_single_setting, fptr.nativePointer, k, val)
}

func (fptr *IFptr) GetSingleSetting(key string) string {
	wideKey, _ := StringToWcharT(key)
	buf := make([]C.wchar_t, 128)
	size := int(C.bridge_libfptr_get_single_setting_func(fptr.functions.libfptr_get_single_setting, fptr.nativePointer, wideKey, &buf[0], C.int(len(buf))))
	if size > len(buf) {
		buf = make([]C.wchar_t, size)
		C.bridge_libfptr_get_single_setting_func(fptr.functions.libfptr_get_single_setting, fptr.nativePointer, wideKey, &buf[0], C.int(len(buf)))
	}
	str, _ := WcharTToString(&buf[0])
	return str
}

func (fptr *IFptr) IsOpened() bool {
	var res = C.bridge_libfptr_is_opened_func(fptr.functions.libfptr_is_opened, fptr.nativePointer)
	if res == 0 {
		return false
	} else {
		return true
	}
}

func (fptr *IFptr) ErrorCode() int {
	return int(C.bridge_libfptr_error_code_func(fptr.functions.libfptr_error_code, fptr.nativePointer))
}

func (fptr *IFptr) ErrorDescription() string {
	buf := make([]C.wchar_t, 128)
	size := int(C.bridge_libfptr_error_description_func(fptr.functions.libfptr_error_description, fptr.nativePointer, &buf[0], C.int(len(buf))))
	if size > len(buf) {
		buf = make([]C.wchar_t, size)
		C.bridge_libfptr_error_description_func(fptr.functions.libfptr_error_description, fptr.nativePointer, &buf[0], C.int(len(buf)))
	}
	str, _ := WcharTToString(&buf[0])

	return str
}

func (fptr *IFptr) ErrorRecommendation() string {
	buf := make([]C.wchar_t, 128)
	size := int(C.bridge_libfptr_error_recommendation_func(fptr.functions.libfptr_error_recommendation, fptr.nativePointer, &buf[0], C.int(len(buf))))
	if size > len(buf) {
		buf = make([]C.wchar_t, size)
		C.bridge_libfptr_error_recommendation_func(fptr.functions.libfptr_error_recommendation, fptr.nativePointer, &buf[0], C.int(len(buf)))
	}
	str, _ := WcharTToString(&buf[0])

	return str
}

func (fptr *IFptr) ResetError() {
	C.bridge_libfptr_reset_error_func(fptr.functions.libfptr_reset_error, fptr.nativePointer)
}

func (fptr *IFptr) SetParam(paramId int32, param interface{}) {
	switch p := param.(type) {
	case int:
		{
			if p < 0 || uint(p) > math.MaxUint32 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int8:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int16:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int32:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint:
		{
			if uint(p) > math.MaxUint32 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint8:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint16:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint32:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case bool:
		{
			boolParam := 0
			if p {
				boolParam = 1
			}
			C.bridge_libfptr_set_param_bool_func(fptr.functions.libfptr_set_param_bool, fptr.nativePointer, C.int(paramId), C.int(boolParam))
		}
	case string:
		{
			buf, _ := StringToWcharT(p)
			C.bridge_libfptr_set_param_str_func(fptr.functions.libfptr_set_param_str, fptr.nativePointer, C.int(paramId), buf)
		}
	case time.Time:
		{
			C.bridge_libfptr_set_param_datetime_func(fptr.functions.libfptr_set_param_datetime, fptr.nativePointer, C.int(paramId),
				C.int(p.Year()), C.int(p.Month()), C.int(p.Day()),
				C.int(p.Hour()), C.int(p.Minute()), C.int(p.Second()))
		}
	case float32:
		{
			C.bridge_libfptr_set_param_double_func(fptr.functions.libfptr_set_param_double, fptr.nativePointer, C.int(paramId), C.double(p))
		}
	case float64:
		{
			C.bridge_libfptr_set_param_double_func(fptr.functions.libfptr_set_param_double, fptr.nativePointer, C.int(paramId), C.double(p))
		}
	case []byte:
		{
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&p[0]), C.int(len(p)))
		}
	case bytes.Buffer:
		{
			bytesParam := p.Bytes()
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&bytesParam[0]), C.int(len(bytesParam)))
		}
	case *bytes.Buffer:
		{
			bytesParam := p.Bytes()
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&bytesParam[0]), C.int(len(bytesParam)))
		}
	default:
		{
			panic(fmt.Sprintf("Invalid parameter type %T", p))
		}
	}
}

func (fptr *IFptr) SetUserParam(paramId int32, param interface{}) {
	switch p := param.(type) {
	case int:
		{
			if p < 0 || uint(p) > math.MaxUint32 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int8:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int16:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int32:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint:
		{
			if uint(p) > math.MaxUint32 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint8:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint16:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint32:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_user_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case bool:
		{
			boolParam := 0
			if p {
				boolParam = 1
			}
			C.bridge_libfptr_set_param_bool_func(fptr.functions.libfptr_set_user_param_bool, fptr.nativePointer, C.int(paramId), C.int(boolParam))
		}
	case string:
		{
			buf, _ := StringToWcharT(p)
			C.bridge_libfptr_set_param_str_func(fptr.functions.libfptr_set_user_param_str, fptr.nativePointer, C.int(paramId), buf)
		}
	case time.Time:
		{
			C.bridge_libfptr_set_param_datetime_func(fptr.functions.libfptr_set_user_param_datetime, fptr.nativePointer, C.int(paramId),
				C.int(p.Year()), C.int(p.Month()), C.int(p.Day()),
				C.int(p.Hour()), C.int(p.Minute()), C.int(p.Second()))
		}
	case float32:
		{
			C.bridge_libfptr_set_param_double_func(fptr.functions.libfptr_set_user_param_double, fptr.nativePointer, C.int(paramId), C.double(p))
		}
	case float64:
		{
			C.bridge_libfptr_set_param_double_func(fptr.functions.libfptr_set_user_param_double, fptr.nativePointer, C.int(paramId), C.double(p))
		}
	case []byte:
		{
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_user_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&p[0]), C.int(len(p)))
		}
	case bytes.Buffer:
		{
			bytesParam := p.Bytes()
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_user_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&bytesParam[0]), C.int(len(bytesParam)))
		}
	case *bytes.Buffer:
		{
			bytesParam := p.Bytes()
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_user_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&bytesParam[0]), C.int(len(bytesParam)))
		}
	default:
		{
			panic(fmt.Sprintf("Invalid parameter type %T", p))
		}
	}
}

func (fptr *IFptr) SetNonPrintableParam(paramId int32, param interface{}) {
	switch p := param.(type) {
	case int:
		{
			if p < 0 || uint(p) > math.MaxUint32 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int8:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int16:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case int32:
		{
			if p < 0 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint:
		{
			if uint(p) > math.MaxUint32 {
				panic(fmt.Sprintf("Invalid parameter value %v", p))
			}
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint8:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint16:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case uint32:
		{
			C.bridge_libfptr_set_param_int_func(fptr.functions.libfptr_set_non_printable_param_int, fptr.nativePointer, C.int(paramId), C.uint(p))
		}
	case bool:
		{
			boolParam := 0
			if p {
				boolParam = 1
			}
			C.bridge_libfptr_set_param_bool_func(fptr.functions.libfptr_set_non_printable_param_bool, fptr.nativePointer, C.int(paramId), C.int(boolParam))
		}
	case string:
		{
			buf, _ := StringToWcharT(p)
			C.bridge_libfptr_set_param_str_func(fptr.functions.libfptr_set_non_printable_param_str, fptr.nativePointer, C.int(paramId), buf)
		}
	case time.Time:
		{
			C.bridge_libfptr_set_param_datetime_func(fptr.functions.libfptr_set_non_printable_param_datetime, fptr.nativePointer, C.int(paramId),
				C.int(p.Year()), C.int(p.Month()), C.int(p.Day()),
				C.int(p.Hour()), C.int(p.Minute()), C.int(p.Second()))
		}
	case float32:
		{
			C.bridge_libfptr_set_param_double_func(fptr.functions.libfptr_set_non_printable_param_double, fptr.nativePointer, C.int(paramId), C.double(p))
		}
	case float64:
		{
			C.bridge_libfptr_set_param_double_func(fptr.functions.libfptr_set_non_printable_param_double, fptr.nativePointer, C.int(paramId), C.double(p))
		}
	case []byte:
		{
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_non_printable_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&p[0]), C.int(len(p)))
		}
	case bytes.Buffer:
		{
			bytesParam := p.Bytes()
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_non_printable_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&bytesParam[0]), C.int(len(bytesParam)))
		}
	case *bytes.Buffer:
		{
			bytesParam := p.Bytes()
			C.bridge_libfptr_set_param_bytearray_func(fptr.functions.libfptr_set_non_printable_param_bytearray, fptr.nativePointer, C.int(paramId),
				(*C.uchar)(&bytesParam[0]), C.int(len(bytesParam)))
		}
	default:
		{
			panic(fmt.Sprintf("Invalid parameter type %T", p))
		}
	}
}

func (fptr *IFptr) GetParamBool(paramId int) bool {
	res := C.bridge_libfptr_get_param_bool_func(fptr.functions.libfptr_get_param_bool, fptr.nativePointer, C.int(paramId))
	if res == 0 {
		return false
	} else {
		return true
	}
}

func (fptr *IFptr) GetParamInt(paramId int) uint {
	return uint(C.bridge_libfptr_get_param_int_func(fptr.functions.libfptr_get_param_int, fptr.nativePointer, C.int(paramId)))
}

func (fptr *IFptr) GetParamDouble(paramId int) float64 {
	return float64(C.bridge_libfptr_get_param_double_func(fptr.functions.libfptr_get_param_double, fptr.nativePointer, C.int(paramId)))
}

func (fptr *IFptr) GetParamString(paramId int) string {
	buf := make([]C.wchar_t, 128)
	size := int(C.bridge_libfptr_get_param_str_func(fptr.functions.libfptr_get_param_str, fptr.nativePointer, C.int(paramId), &buf[0], C.int(len(buf))))
	if size > len(buf) {
		buf = make([]C.wchar_t, size)
		C.bridge_libfptr_get_param_str_func(fptr.functions.libfptr_get_param_str, fptr.nativePointer, C.int(paramId), &buf[0], C.int(len(buf)))
	}
	str, _ := WcharTToString(&buf[0])
	return str
}

func (fptr *IFptr) GetParamByteArray(paramId int) []byte {
	buf := make([]byte, 128)
	size := int(C.bridge_libfptr_get_param_bytearray_func(fptr.functions.libfptr_get_param_bytearray, fptr.nativePointer, C.int(paramId), (*C.uchar)(&buf[0]), C.int(len(buf))))
	if size > len(buf) {
		buf = make([]byte, size)
		C.bridge_libfptr_get_param_bytearray_func(fptr.functions.libfptr_get_param_bytearray, fptr.nativePointer, C.int(paramId), (*C.uchar)(&buf[0]), C.int(len(buf)))
	}

	return buf[:size]
}

func (fptr *IFptr) GetParamDateTime(paramId int) time.Time {
	var year, month, day, hour, minute, second C.int
	C.bridge_libfptr_get_param_datetime_func(fptr.functions.libfptr_get_param_datetime, fptr.nativePointer, C.int(paramId), &year, &month, &day, &hour, &minute, &second)
	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.Now().Location())
}

func (fptr *IFptr) IsParamAvailable(paramId int) bool {
	res := C.bridge_libfptr_is_param_available_func(fptr.functions.libfptr_is_param_available, fptr.nativePointer, C.int(paramId))
	if res > 0 {
		return true
	} else {
		return false
	}
}

func (fptr *IFptr) LogWrite(tag string, level int, message string) error {
	tg, _ := StringToWcharT(tag)
	msg, _ := StringToWcharT(message)
	err := C.bridge_libfptr_log_write_func(fptr.functions.libfptr_log_write, fptr.nativePointer, tg, C.int(level), msg)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ChangeLabel(label string) error {
	str, _ := StringToWcharT(label)
	err := C.bridge_libfptr_change_label_func(fptr.functions.libfptr_change_label, fptr.nativePointer, str)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ShowProperties(parent_type int, parent unsafe.Pointer) error {
	err := C.bridge_libfptr_show_properties_func(fptr.functions.libfptr_show_properties, fptr.nativePointer, C.uint(parent_type), parent)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ApplySingleSettings() error {
	if fptr.functions.libfptr_apply_single_settings == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_apply_single_settings\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_apply_single_settings, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) Open() error {
	if fptr.functions.libfptr_open == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_open\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_open, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) Close() error {
	if fptr.functions.libfptr_close == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_close\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_close, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ResetParams() error {
	if fptr.functions.libfptr_reset_params == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_reset_params\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_reset_params, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) RunCommand() error {
	if fptr.functions.libfptr_run_command == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_run_command\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_run_command, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) Beep() error {
	if fptr.functions.libfptr_beep == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_beep\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_beep, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) OpenDrawer() error {
	if fptr.functions.libfptr_open_drawer == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_open_drawer\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_open_drawer, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) Cut() error {
	if fptr.functions.libfptr_cut == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_cut\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_cut, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) DevicePoweroff() error {
	if fptr.functions.libfptr_device_poweroff == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_device_poweroff\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_device_poweroff, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) DeviceReboot() error {
	if fptr.functions.libfptr_device_reboot == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_device_reboot\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_device_reboot, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) OpenShift() error {
	if fptr.functions.libfptr_open_shift == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_open_shift\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_open_shift, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ResetSummary() error {
	if fptr.functions.libfptr_reset_summary == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_reset_summary\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_reset_summary, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) InitDevice() error {
	if fptr.functions.libfptr_init_device == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_init_device\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_init_device, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) QueryData() error {
	if fptr.functions.libfptr_query_data == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_query_data\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_query_data, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CashIncome() error {
	if fptr.functions.libfptr_cash_income == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_cash_income\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_cash_income, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CashOutcome() error {
	if fptr.functions.libfptr_cash_outcome == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_cash_outcome\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_cash_outcome, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) OpenReceipt() error {
	if fptr.functions.libfptr_open_receipt == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_open_receipt\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_open_receipt, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CancelReceipt() error {
	if fptr.functions.libfptr_cancel_receipt == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_cancel_receipt\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_cancel_receipt, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CloseReceipt() error {
	if fptr.functions.libfptr_close_receipt == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_close_receipt\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_close_receipt, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CheckDocumentClosed() error {
	if fptr.functions.libfptr_check_document_closed == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_check_document_closed\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_check_document_closed, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReceiptTotal() error {
	if fptr.functions.libfptr_receipt_total == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_receipt_total\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_receipt_total, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReceiptTax() error {
	if fptr.functions.libfptr_receipt_tax == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_receipt_tax\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_receipt_tax, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) Registration() error {
	if fptr.functions.libfptr_registration == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_registration\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_registration, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) Payment() error {
	if fptr.functions.libfptr_payment == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_payment\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_payment, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) Report() error {
	if fptr.functions.libfptr_report == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_report\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_report, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) PrintText() error {
	if fptr.functions.libfptr_print_text == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_print_text\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_print_text, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) PrintCliche() error {
	if fptr.functions.libfptr_print_cliche == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_print_cliche\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_print_cliche, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) BeginNonfiscalDocument() error {
	if fptr.functions.libfptr_begin_nonfiscal_document == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_begin_nonfiscal_document\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_begin_nonfiscal_document, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) EndNonfiscalDocument() error {
	if fptr.functions.libfptr_end_nonfiscal_document == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_end_nonfiscal_document\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_end_nonfiscal_document, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) PrintBarcode() error {
	if fptr.functions.libfptr_print_barcode == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_print_barcode\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_print_barcode, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) PrintPicture() error {
	if fptr.functions.libfptr_print_picture == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_print_picture\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_print_picture, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) PrintPictureByNumber() error {
	if fptr.functions.libfptr_print_picture_by_number == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_print_picture_by_number\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_print_picture_by_number, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UploadPictureFromFile() error {
	if fptr.functions.libfptr_upload_picture_from_file == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_upload_picture_from_file\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_upload_picture_from_file, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ClearPictures() error {
	if fptr.functions.libfptr_clear_pictures == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_clear_pictures\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_clear_pictures, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) WriteDeviceSettingRaw() error {
	if fptr.functions.libfptr_write_device_setting_raw == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_write_device_setting_raw\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_write_device_setting_raw, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReadDeviceSettingRaw() error {
	if fptr.functions.libfptr_read_device_setting_raw == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_read_device_setting_raw\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_read_device_setting_raw, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CommitSettings() error {
	if fptr.functions.libfptr_commit_settings == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_commit_settings\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_commit_settings, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) InitSettings() error {
	if fptr.functions.libfptr_init_settings == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_init_settings\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_init_settings, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ResetSettings() error {
	if fptr.functions.libfptr_reset_settings == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_reset_settings\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_reset_settings, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) WriteDateTime() error {
	if fptr.functions.libfptr_write_date_time == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_write_date_time\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_write_date_time, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) WriteLicense() error {
	if fptr.functions.libfptr_write_license == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_write_license\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_write_license, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) FnOperation() error {
	if fptr.functions.libfptr_fn_operation == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_fn_operation\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_fn_operation, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) FnQueryData() error {
	if fptr.functions.libfptr_fn_query_data == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_fn_query_data\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_fn_query_data, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) FnWriteAttributes() error {
	if fptr.functions.libfptr_fn_write_attributes == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_fn_write_attributes\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_fn_write_attributes, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ExternalDevicePowerOn() error {
	if fptr.functions.libfptr_external_device_power_on == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_external_device_power_on\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_external_device_power_on, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ExternalDevicePowerOff() error {
	if fptr.functions.libfptr_external_device_power_off == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_external_device_power_off\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_external_device_power_off, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ExternalDeviceWriteData() error {
	if fptr.functions.libfptr_external_device_write_data == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_external_device_write_data\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_external_device_write_data, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ExternalDeviceReadData() error {
	if fptr.functions.libfptr_external_device_read_data == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_external_device_read_data\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_external_device_read_data, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) OperatorLogin() error {
	if fptr.functions.libfptr_operator_login == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_operator_login\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_operator_login, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ProcessJson() error {
	if fptr.functions.libfptr_process_json == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_process_json\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_process_json, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReadDeviceSetting() error {
	if fptr.functions.libfptr_read_device_setting == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_read_device_setting\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_read_device_setting, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) WriteDeviceSetting() error {
	if fptr.functions.libfptr_write_device_setting == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_write_device_setting\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_write_device_setting, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) BeginReadRecords() error {
	if fptr.functions.libfptr_begin_read_records == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_begin_read_records\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_begin_read_records, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReadNextRecord() error {
	if fptr.functions.libfptr_read_next_record == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_read_next_record\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_read_next_record, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) EndReadRecords() error {
	if fptr.functions.libfptr_end_read_records == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_end_read_records\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_end_read_records, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UserMemoryOperation() error {
	if fptr.functions.libfptr_user_memory_operation == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_user_memory_operation\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_user_memory_operation, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ContinuePrint() error {
	if fptr.functions.libfptr_continue_print == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_continue_print\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_continue_print, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) InitMgm() error {
	if fptr.functions.libfptr_init_mgm == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_init_mgm\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_init_mgm, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UtilFormTlv() error {
	if fptr.functions.libfptr_util_form_tlv == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_util_form_tlv\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_util_form_tlv, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UtilFormNomenclature() error {
	if fptr.functions.libfptr_util_form_nomenclature == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_util_form_nomenclature\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_util_form_nomenclature, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UtilMapping() error {
	if fptr.functions.libfptr_util_mapping == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_util_mapping\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_util_mapping, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReadModelFlags() error {
	if fptr.functions.libfptr_read_model_flags == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_read_model_flags\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_read_model_flags, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) LineFeed() error {
	if fptr.functions.libfptr_line_feed == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_line_feed\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_line_feed, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) FlashFirmware() error {
	if fptr.functions.libfptr_flash_firmware == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_flash_firmware\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_flash_firmware, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) SoftLockInit() error {
	if fptr.functions.libfptr_soft_lock_init == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_soft_lock_init\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_soft_lock_init, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) SoftLockQuerySessionCode() error {
	if fptr.functions.libfptr_soft_lock_query_session_code == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_soft_lock_query_session_code\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_soft_lock_query_session_code, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) SoftLockValidate() error {
	if fptr.functions.libfptr_soft_lock_validate == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_soft_lock_validate\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_soft_lock_validate, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UtilCalcTax() error {
	if fptr.functions.libfptr_util_calc_tax == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_util_calc_tax\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_util_calc_tax, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) DownloadPicture() error {
	if fptr.functions.libfptr_download_picture == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_download_picture\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_download_picture, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) BluetoothRemovePairedDevices() error {
	if fptr.functions.libfptr_bluetooth_remove_paired_devices == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_bluetooth_remove_paired_devices\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_bluetooth_remove_paired_devices, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UtilTagInfo() error {
	if fptr.functions.libfptr_util_tag_info == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_util_tag_info\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_util_tag_info, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UtilContainerVersions() error {
	if fptr.functions.libfptr_util_container_versions == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_util_container_versions\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_util_container_versions, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ActivateLicenses() error {
	if fptr.functions.libfptr_activate_licenses == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_activate_licenses\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_activate_licenses, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) RemoveLicenses() error {
	if fptr.functions.libfptr_remove_licenses == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_remove_licenses\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_remove_licenses, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) EnterKeys() error {
	if fptr.functions.libfptr_enter_keys == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_enter_keys\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_enter_keys, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ValidateKeys() error {
	if fptr.functions.libfptr_validate_keys == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_validate_keys\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_validate_keys, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) EnterSerialNumber() error {
	if fptr.functions.libfptr_enter_serial_number == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_enter_serial_number\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_enter_serial_number, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) GetSerialNumberRequest() error {
	if fptr.functions.libfptr_get_serial_number_request == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_get_serial_number_request\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_get_serial_number_request, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UploadPixelBuffer() error {
	if fptr.functions.libfptr_upload_pixel_buffer == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_upload_pixel_buffer\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_upload_pixel_buffer, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) DownloadPixelBuffer() error {
	if fptr.functions.libfptr_download_pixel_buffer == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_download_pixel_buffer\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_download_pixel_buffer, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) PrintPixelBuffer() error {
	if fptr.functions.libfptr_print_pixel_buffer == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_print_pixel_buffer\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_print_pixel_buffer, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UtilConvertTagValue() error {
	if fptr.functions.libfptr_util_convert_tag_value == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_util_convert_tag_value\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_util_convert_tag_value, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ParseMarkingCode() error {
	if fptr.functions.libfptr_parse_marking_code == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_parse_marking_code\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_parse_marking_code, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CallScript() error {
	if fptr.functions.libfptr_call_script == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_call_script\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_call_script, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) SetHeaderLines() error {
	if fptr.functions.libfptr_set_header_lines == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_set_header_lines\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_set_header_lines, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) SetFooterLines() error {
	if fptr.functions.libfptr_set_footer_lines == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_set_footer_lines\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_set_footer_lines, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UploadPictureCliche() error {
	if fptr.functions.libfptr_upload_picture_cliche == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_upload_picture_cliche\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_upload_picture_cliche, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UploadPictureMemory() error {
	if fptr.functions.libfptr_upload_picture_memory == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_upload_picture_memory\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_upload_picture_memory, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UploadPixelBufferCliche() error {
	if fptr.functions.libfptr_upload_pixel_buffer_cliche == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_upload_pixel_buffer_cliche\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_upload_pixel_buffer_cliche, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UploadPixelBufferMemory() error {
	if fptr.functions.libfptr_upload_pixel_buffer_memory == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_upload_pixel_buffer_memory\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_upload_pixel_buffer_memory, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ExecDriverScript() error {
	if fptr.functions.libfptr_exec_driver_script == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_exec_driver_script\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_exec_driver_script, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UploadDriverScript() error {
	if fptr.functions.libfptr_upload_driver_script == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_upload_driver_script\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_upload_driver_script, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ExecDriverScriptById() error {
	if fptr.functions.libfptr_exec_driver_script_by_id == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_exec_driver_script_by_id\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_exec_driver_script_by_id, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) WriteUniversalCountersSettings() error {
	if fptr.functions.libfptr_write_universal_counters_settings == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_write_universal_counters_settings\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_write_universal_counters_settings, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReadUniversalCountersSettings() error {
	if fptr.functions.libfptr_read_universal_counters_settings == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_read_universal_counters_settings\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_read_universal_counters_settings, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) QueryUniversalCountersState() error {
	if fptr.functions.libfptr_query_universal_counters_state == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_query_universal_counters_state\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_query_universal_counters_state, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ResetUniversalCounters() error {
	if fptr.functions.libfptr_reset_universal_counters == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_reset_universal_counters\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_reset_universal_counters, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CacheUniversalCounters() error {
	if fptr.functions.libfptr_cache_universal_counters == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_cache_universal_counters\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_cache_universal_counters, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReadUniversalCounterSum() error {
	if fptr.functions.libfptr_read_universal_counter_sum == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_read_universal_counter_sum\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_read_universal_counter_sum, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReadUniversalCounterQuantity() error {
	if fptr.functions.libfptr_read_universal_counter_quantity == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_read_universal_counter_quantity\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_read_universal_counter_quantity, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ClearUniversalCountersCache() error {
	if fptr.functions.libfptr_clear_universal_counters_cache == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_clear_universal_counters_cache\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_clear_universal_counters_cache, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) DisableOfdChannel() error {
	if fptr.functions.libfptr_disable_ofd_channel == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_disable_ofd_channel\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_disable_ofd_channel, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) EnableOfdChannel() error {
	if fptr.functions.libfptr_enable_ofd_channel == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_enable_ofd_channel\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_enable_ofd_channel, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ValidateJson() error {
	if fptr.functions.libfptr_validate_json == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_validate_json\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_validate_json, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ReflectionCall() error {
	if fptr.functions.libfptr_reflection_call == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_reflection_call\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_reflection_call, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) GetRemoteServerInfo() error {
	if fptr.functions.libfptr_get_remote_server_info == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_get_remote_server_info\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_get_remote_server_info, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) BeginMarkingCodeValidation() error {
	if fptr.functions.libfptr_begin_marking_code_validation == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_begin_marking_code_validation\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_begin_marking_code_validation, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CancelMarkingCodeValidation() error {
	if fptr.functions.libfptr_cancel_marking_code_validation == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_cancel_marking_code_validation\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_cancel_marking_code_validation, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) GetMarkingCodeValidationStatus() error {
	if fptr.functions.libfptr_get_marking_code_validation_status == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_get_marking_code_validation_status\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_get_marking_code_validation_status, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) AcceptMarkingCode() error {
	if fptr.functions.libfptr_accept_marking_code == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_accept_marking_code\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_accept_marking_code, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) DeclineMarkingCode() error {
	if fptr.functions.libfptr_decline_marking_code == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_decline_marking_code\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_decline_marking_code, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) UpdateFnmKeys() error {
	if fptr.functions.libfptr_update_fnm_keys == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_update_fnm_keys\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_update_fnm_keys, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) WriteSalesNotice() error {
	if fptr.functions.libfptr_write_sales_notice == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_write_sales_notice\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_write_sales_notice, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) CheckMarkingCodeValidationsReady() error {
	if fptr.functions.libfptr_check_marking_code_validations_ready == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_check_marking_code_validations_ready\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_check_marking_code_validations_ready, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) ClearMarkingCodeValidationResult() error {
	if fptr.functions.libfptr_clear_marking_code_validation_result == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_clear_marking_code_validation_result\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_clear_marking_code_validation_result, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) PingMarkingServer() error {
	if fptr.functions.libfptr_ping_marking_server == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_ping_marking_server\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_ping_marking_server, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) GetMarkingServerStatus() error {
	if fptr.functions.libfptr_get_marking_server_status == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_get_marking_server_status\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_get_marking_server_status, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) IsDriverLocked() error {
	if fptr.functions.libfptr_is_driver_locked == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_is_driver_locked\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_is_driver_locked, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) GetLastDocumentJournal() error {
	if fptr.functions.libfptr_get_last_document_journal == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_get_last_document_journal\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_get_last_document_journal, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}

func (fptr *IFptr) FindDocumentInJournal() error {
	if fptr.functions.libfptr_find_document_in_journal == nil {
		return &Error{LIBFPTR_ERROR_NOT_LOADED, fmt.Sprintf("Can`t find method \"libfptr_find_document_in_journal\" in library")}
	}
	err := C.bridge_libfptr_simple_func(fptr.functions.libfptr_find_document_in_journal, fptr.nativePointer)
	if err == 0 {
		return nil
	}
	return &Error{fptr.ErrorCode(), fptr.ErrorDescription()}
}
