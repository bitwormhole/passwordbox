package com.bitwormhole.passwordbox.app.ui;

import android.os.Bundle;

import androidx.annotation.Nullable;

import com.bitwormhole.passwordbox.app.R;
import com.bitwormhole.passwordbox.app.ui.tools.ActivityHelper;

public class DebugHomeActivity extends BaseActivity {

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.layout_debug_home);


        ActivityHelper ah = new ActivityHelper(this);
        ah.setupButtonToOpenActivity(R.id.button_debug_promise, DebugPromiseActivity.class);
        ah.setupButtonToOpenActivity(R.id.button_debug_root_key, DebugRootKeyActivity.class);

    }

}
