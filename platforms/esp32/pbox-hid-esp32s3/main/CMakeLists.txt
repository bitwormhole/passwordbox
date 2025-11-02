idf_component_register(
    SRCS "pbox_array.c" "pbox_app_context.c" "module_auto_test.c" "module_ble.c" "module_sd_card.c" "module_usb_hid.c" "pbox_errors.c" "pbox_module.c" "pbox_logger.c" "module_wifi.c" "pbox_app.c" "tusb_hid_example_main.c"
    INCLUDE_DIRS "."
    PRIV_REQUIRES esp_driver_gpio
    )
