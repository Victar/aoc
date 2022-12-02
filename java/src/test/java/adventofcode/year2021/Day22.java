package adventofcode.year2021;

import java.util.ArrayList;
import java.util.List;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day22 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day22/input.txt");
		final List<ActionCube> actions = new ArrayList<>();
		for (final String input : data) {
			actions.add(new ActionCube(input));
		}
		long startX = Long.MAX_VALUE;
		long endX = Long.MIN_VALUE;
		long startY = Long.MAX_VALUE;
		long endY = Long.MIN_VALUE;
		long startZ = Long.MAX_VALUE;
		long endZ = Long.MIN_VALUE;
		for (final ActionCube action : actions) {
			System.out.println(action);
			startX = Math.min(startX, action.startX);
			startY = Math.min(startY, action.startY);
			startZ = Math.min(startZ, action.startZ);
			endX = Math.max(endX, action.endX);
			endY = Math.max(endY, action.endY);
			endZ = Math.max(endZ, action.endZ);
		}
		int count = 0;

		for (int x = -50; x <= 50; x++) {
			for (int y = -50; y <= 50; y++) {
				for (int z = -50; z <= 50; z++) {
					if (isLight(actions, x, y, z)) {
						count++;
					}
				}
			}
		}
		System.out.println(count);
	}

	public boolean isLight(final List<ActionCube> actions, final int x, final int y, final int z) {
		boolean result = false;
		for (final ActionCube action : actions) {
			result = action.check(x, y, z, result);
		}
		return result;
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day22/input.txt");
		final List<ActionCube> actions = new ArrayList<>();
		final List<Cube> cubes = new ArrayList<>();
		for (final String input : data) {
			actions.add(new ActionCube(input));
		}
		for (final ActionCube action : actions) {
			cubes.addAll(processAction(action, cubes));
		}

		long count = 0;
		for (final Cube cube : cubes) {
			count += cube.countSpace();
		}
		System.out.println(count);
	}

	List<Cube> processAction(final ActionCube action, final List<Cube> cubes) {
		final List<Cube> current = new ArrayList<>();
		for (final Cube cube : cubes) {
			final Cube inter = cube.intersect(action);
			if (inter != null) {
				current.add(inter);
			}
		}
		if (action.isOn) {
			current.add(action);
		}
		return current;
	}

	@Data class Cube {

		long startX;
		long endX;
		long startY;
		long endY;
		long startZ;
		long endZ;
		boolean isOn;

		public Cube() {

		}

		public Cube(final long startX, final long endX, final long startY, final long endY, final long startZ, final long endZ,
		            final boolean isOn) {
			this.startX = startX;
			this.endX = endX;
			this.startY = startY;
			this.endY = endY;
			this.startZ = startZ;
			this.endZ = endZ;
			this.isOn = isOn;
		}

		public Cube intersect(final Cube compare) {
			if (this.startX > compare.endX || this.endX < compare.startX //x
					|| this.startY > compare.endY || this.endY < compare.startY //x
					|| this.startZ > compare.endZ || this.endZ < compare.startZ) {
				return null;
			} else {
				return new Cube(Math.max(this.startX, compare.startX), Math.min(this.endX, compare.endX), //x
						Math.max(this.startY, compare.startY), Math.min(this.endY, compare.endY), //y
						Math.max(this.startZ, compare.startZ), Math.min(this.endZ, compare.endZ), !this.isOn);
			}
		}

		public long countSpace() {
			return (this.endX - this.startX + 1) * (this.endY - this.startY + 1) * (this.endZ - this.startZ + 1) * (this.isOn ? 1 : -1);
		}

		@Override public String toString() {
			return (this.isOn ? "on" : "off") + " x=" + this.startX + ".." + this.endX + ",y=" + this.startY + ".." + this.endY + ",z="
					+ this.startZ + ".." + this.endZ + "  count:" + this.countSpace();
		}
	}

	@Data class ActionCube extends Cube {

		public ActionCube(final String input) {
			final String[] arr = StringUtils.split(input, " ");
			this.isOn = "on".equals(arr[0]);
			final String[] arr2 = StringUtils.split(arr[1], ",");
			//read
			final String xInput = arr2[0].substring(2).replace("..", " ");
			final String[] arrx = StringUtils.split(xInput, " ");
			this.startX = Integer.parseInt(arrx[0]);
			this.endX = Integer.parseInt(arrx[1]);

			final String yInput = arr2[1].substring(2).replace("..", " ");
			final String[] arry = StringUtils.split(yInput, " ");
			this.startY = Integer.parseInt(arry[0]);
			this.endY = Integer.parseInt(arry[1]);

			final String zInput = arr2[2].substring(2).replace("..", " ");
			final String[] arrz = StringUtils.split(zInput, " ");
			this.startZ = Integer.parseInt(arrz[0]);
			this.endZ = Integer.parseInt(arrz[1]);
		}

		boolean check(final int x, final int y, final int z, final boolean current) {
			if (this.startX <= x && x <= this.endX && this.startY <= y && y <= this.endY && this.startZ <= z && z <= this.endZ) {
				return this.isOn;
			} else {
				return current;
			}

		}

		@Override public String toString() {
			return super.toString();
		}

	}

}
