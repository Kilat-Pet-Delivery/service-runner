CREATE TABLE crate_specs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    runner_id UUID NOT NULL REFERENCES runners(id) ON DELETE CASCADE,
    size VARCHAR(20) NOT NULL,
    pet_types JSONB NOT NULL DEFAULT '[]',
    max_weight_kg DECIMAL(6,2) NOT NULL,
    width_cm DECIMAL(6,2),
    height_cm DECIMAL(6,2),
    depth_cm DECIMAL(6,2),
    ventilated BOOLEAN DEFAULT FALSE,
    temperature_controlled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_crate_specs_runner ON crate_specs(runner_id);
