package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Date;
import java.util.HashSet;
import java.util.Set;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;
import lombok.ToString;

public class Day17 extends BaseTest {

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day17/input.txt");
		final int end = data.size() + 1;
		final int start = -1;
		Set<Cord> active = new HashSet<>();
		for (int i = 0; i < data.size(); i++) {
			for (int j = 0; j < data.get(i).length(); j++) {
				if (data.get(i).charAt(j) == '#') active.add(Cord.of(i, j, 0, 0));
			}
		}
		long time = System.currentTimeMillis();
		for (int i = 0; i < 6; i++) {
			active = doCheck(active, start - i, end + i);
			final long currentTimeMillis = System.currentTimeMillis();
			System.out.println(active.size() + "     time: " + new Date() + " diff: " + (currentTimeMillis - time));
			time = currentTimeMillis;
		}
	}

	public Set<Cord> doCheck(final Set<Cord> activeSet, final int start, final int end) {
		System.out.println("start: " + start + " end: " + end);
		final HashSet<Cord> newState = new HashSet<>();
		for (int i = start; i < end; i++) {
			for (int j = start; j < end; j++) {
				for (int k = start; k < end; k++) {
					for (int w = start; w < end; w++) {
						final Cord toCheck = Cord.of(i, j, k, w);
						final int nbr = checkNeighbors(toCheck, activeSet);
						final boolean contains = activeSet.contains(toCheck);
						if (!contains && nbr == 3) {
							newState.add(toCheck);
						}
						if (contains && (nbr == 3 || nbr == 2)) {
							newState.add(toCheck);
						}
					}
				}
			}
		}
		return newState;
	}

	public int checkNeighbors(final Cord cord, final Set<Cord> cubesDimension) {
		final int[] dX = { -1, 0, 1 };
		final int[] dY = { -1, 0, 1 };
		final int[] dZ = { -1, 0, 1 };
		final int[] dW = { -1, 0, 1 };

		int active = 0;

		for (int i = 0; i < dX.length; i++) {
			for (int j = 0; j < dY.length; j++) {
				for (int k = 0; k < dZ.length; k++) {
					for (int w = 0; w < dW.length; w++) {
						if (i != 1 || j != 1 || k != 1 || w != 1) {
							final Cord neighborToCheck = Cord.of(cord.x + dX[i], cord.y + dY[j], cord.z + dZ[k], cord.w + dZ[w]);
							//							Cord neighbor = cubesDimension.stream().filter(c -> c.equals(neighborToCheck)).findAny().orElse(null);
							//							if (neighbor != null) {
							//								active++;
							//							}
							if (cubesDimension.contains(neighborToCheck)) {
								active++;
							}
						}
					}
				}
			}
		}
		return active;
	}

	@Data @ToString static class Cord {

		int x;
		int y;
		int z;
		int w;

		public Cord(final int x, final int y, final int z, final int w) {
			this.x = x;
			this.y = y;
			this.z = z;
			this.w = w;

		}

		public static Cord of(final int x, final int y, final int z, final int w) {
			return new Cord(x, y, z, w);
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Cord cord = (Cord) o;

			if (this.x != cord.x) return false;
			if (this.y != cord.y) return false;
			if (this.z != cord.z) return false;
			return this.w == cord.w;
		}

		@Override public int hashCode() {
			int result = this.x;
			result = 31 * result + this.y;
			result = 31 * result + this.z;
			result = 31 * result + this.w;
			return result;
		}
	}
}
